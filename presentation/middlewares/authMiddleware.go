package middlewares

import (
	"fmt"
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// TokenExtractor é um middleware que extrai o token JWT de diferentes fontes na requisição
// seguindo uma ordem de prioridade: cookies, header de autorização e por fim query parameters.
// Uma vez encontrado, o token é adicionado ao header Authorization para processamento pelos
// middlewares subsequentes.
func TokenExtractor() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[LOG] TokenExtractor chamado para rota: %s\n", c.Request.URL.Path)
		// Só executa se o header Authorization ainda não existe
		if c.GetHeader("Authorization") == "" {
			var token string

			// 1. Tenta obter do cookie - primeira fonte de prioridade
			token, _ = c.Cookie("jwt")
			if token == "" {
				token, _ = c.Cookie("token") // fallback para cookie alternativo
			}

			// 2. Tenta obter de query parameter - última fonte de prioridade
			if token == "" {
				token = c.Query("token")
			}

			// Só adiciona o header se o token realmente existir e não for vazio
			if token != "" && len(token) > 10 { // JWTs válidos costumam ser maiores que 10 caracteres
				c.Request.Header.Set("Authorization", "Bearer "+token)
				fmt.Printf("TokenExtractor: Token encontrado\n")
			} else {
				fmt.Printf("TokenExtractor: Nenhum token encontrado ou token muito curto\n")
			}
		}
		// Continua a execução da cadeia de middlewares
		c.Next()
	}
}

// TokenDebugger é um middleware para fazer debug das claims do JWT
func TokenDebugger() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		fmt.Printf("==== DEBUG TOKEN JWT ====\n")
		fmt.Printf("Rota: %s %s\n", c.Request.Method, c.Request.URL.Path)
		fmt.Printf("Claims completas do token:\n")
		for key, value := range claims {
			fmt.Printf("  %s: %v (tipo: %T)\n", key, value, value)
		}
		fmt.Printf("========================\n")

		c.Next()
	}
}

// UserIDExtractor é um middleware que extrai o ID do usuário das claims do JWT
// e o disponibiliza no contexto como "user_id_logged"
func UserIDExtractor() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		// Tenta extrair o ID do usuário das claims
		if userID, exists := claims["id"]; exists {
			// Converte o userID para string se necessário
			var userIDStr string
			switch v := userID.(type) {
			case float64:
				userIDStr = fmt.Sprintf("%.0f", v)
			case string:
				userIDStr = v
			case int:
				userIDStr = fmt.Sprintf("%d", v)
			default:
				userIDStr = fmt.Sprintf("%v", v)
			}

			// Define o user_id_logged no contexto
			c.Set("user_id_logged", userIDStr)
			fmt.Printf("[UserIDExtractor] ID do usuário extraído: %s\n", userIDStr)
		} else {
			fmt.Printf("[UserIDExtractor] AVISO: ID do usuário não encontrado nas claims\n")
		}

		c.Next()
	}
}

// AdminRequired é um middleware que verifica se o usuário autenticado possui privilégios de administrador
// através da propriedade "is_admin" nas claims do JWT. Caso contrário, a requisição é abortada com
// status 403 Forbidden.
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrai as claims do token JWT da requisição atual
		claims := jwt.ExtractClaims(c)

		fmt.Printf("==== VERIFICAÇÃO ADMIN ====\n")
		fmt.Printf("Rota: %s %s\n", c.Request.Method, c.Request.URL.Path)
		fmt.Printf("Claims recebidas: %+v\n", claims)

		// Verifica se a claim "is_admin" existe e se seu valor é true
		isAdmin, exists := claims["is_admin"]
		fmt.Printf("is_admin existe: %v, valor: %v (tipo: %T)\n", exists, isAdmin, isAdmin)

		if !exists || isAdmin != true {
			fmt.Printf("Acesso negado - usuário não é admin\n")
			fmt.Printf("==========================\n")
			// Retorna erro 403 para usuários não administradores
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Acesso restrito a administradores",
			})
			c.Abort() // Interrompe a execução de middlewares subsequentes
			return
		}

		fmt.Printf("Acesso permitido - usuário é admin\n")
		fmt.Printf("==========================\n")
		// Se for admin, permite a continuação da requisição
		c.Next()
	}
}

// login define a estrutura esperada para as requisições de autenticação
// contendo email e senha do usuário. As tags de binding garantem validação
// básica dos campos.
type login struct {
	Email    string `json:"email" binding:"required,email"` // Email validado pelo formato
	Password string `json:"password" binding:"required"`    // Senha obrigatória
}

// SetupJWTMiddleware configura e retorna uma instância do middleware JWT para autenticação
// Recebe uma instância do serviço de usuário e o segredo JWT.
func SetupJWTMiddleware(userService *services.UserService, jwtSecret string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "library-api",      // Nome do domínio de autenticação
		Key:         []byte(jwtSecret),  // Chave secreta para assinatura dos tokens
		Timeout:     time.Hour * 24,     // Duração de validade do token: 24 horas
		MaxRefresh:  time.Hour * 24 * 7, // Período máximo em que o token pode ser renovado: 7 dias
		IdentityKey: "id",               // Chave que identifica o usuário nas claims

		// Configurações de cookies de autenticação
		SendCookie:     true,                     // Envia token como cookie
		CookieName:     "jwt",                    // Nome do cookie
		CookieMaxAge:   24 * time.Hour,           // Tempo de vida do cookie
		CookieDomain:   "",                       // Domínio do cookie (vazio = domínio atual)
		SecureCookie:   false,                    // Cookie não requer HTTPS (alterar para true em produção)
		CookieHTTPOnly: true,                     // Cookie não acessível via JavaScript (proteção XSS)
		CookieSameSite: http.SameSiteDefaultMode, // Política de SameSite do cookie

		// Configuração de lookup do token
		TokenLookup:   "cookie:jwt,header:Authorization", // Ordem de busca do token
		TokenHeadName: "Bearer",                          // Prefixo esperado no header
		TimeFunc:      time.Now,                          // Função para obter a hora atual

		// Authenticator: função responsável por validar as credenciais do usuário
		// e retornar os dados do usuário para geração do token JWT
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			// Lê e valida os dados de login do corpo da requisição
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			fmt.Printf("Login - Tentativa para email: %s\n", loginVals.Email)

			// Autentica o usuário usando o serviço
			user, err := userService.AuthenticateUser(loginVals.Email, loginVals.Password)
			if err != nil {
				fmt.Printf("Login - Falha: %v\n", err)
				return nil, jwt.ErrFailedAuthentication
			}

			fmt.Printf("Login - Sucesso para usuário: %s (ID: %d)\n", user.Email, user.ID)
			fmt.Printf("Tipo do usuário retornado: %T\n", user)

			// Cria um DTO com os dados necessários para o token JWT
			userDTO := &dtos.UserResponseDTO{
				UserID:    user.ID,
				UserName:  user.Name,
				UserEmail: user.Email,
				Admin:     user.Admin,
			}

			fmt.Printf("Criando token com user info: ID=%d, Email=%s\n",
				userDTO.UserID, userDTO.UserEmail)

			// Salva o objeto user no contexto para ser usado no LoginResponse
			c.Set("user", userDTO)

			return userDTO, nil
		},

		// PayloadFunc: função que define quais dados do usuário serão incluídos no token JWT
		// como claims personalizadas
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Printf("PayloadFunc recebeu dados do tipo: %T\n", data)

			// Tenta converter o objeto para o tipo esperado
			if user, ok := data.(*dtos.UserResponseDTO); ok {
				fmt.Printf("Convertido com sucesso para UserResponseDTO: ID=%d, Email=%s\n",
					user.UserID, user.UserEmail)

				// Define as claims que serão adicionadas ao token JWT
				return jwt.MapClaims{
					"id":       user.UserID,
					"email":    user.UserEmail,
					"is_admin": user.Admin == 1,
				}
			}

			// Tratamento de erro para tipo inesperado
			fmt.Printf("AVISO: Falha na conversão para UserResponseDTO. Tentando alternativas...\n")
			fmt.Printf("Conteúdo do objeto recebido: %+v\n", data)

			// Retorna claims vazias se a conversão falhar
			return jwt.MapClaims{}
		},
		// IdentityHandler: extrai a identidade do usuário das claims do JWT
		// durante a validação do token
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			fmt.Printf("==== IDENTITY HANDLER ====\n")
			fmt.Printf("Rota: %s %s\n", c.Request.Method, c.Request.URL.Path)
			fmt.Printf("Claims completas: %+v\n", claims)

			// Verifica se todas as claims necessárias estão presentes
			idVal, idExists := claims["id"]
			emailVal, emailExists := claims["email"]
			adminVal, adminExists := claims["is_admin"]

			fmt.Printf("Claim 'id' existe: %v, valor: %v (tipo: %T)\n", idExists, idVal, idVal)
			fmt.Printf("Claim 'email' existe: %v, valor: %v (tipo: %T)\n", emailExists, emailVal, emailVal)
			fmt.Printf("Claim 'is_admin' existe: %v, valor: %v (tipo: %T)\n", adminExists, adminVal, adminVal)

			if !idExists || !emailExists {
				fmt.Printf("ERRO: JWT incompleto ou inválido - claims obrigatórias ausentes\n")
				fmt.Printf("========================\n")
				return nil
			}

			// Conversão segura do ID para o tipo correto
			var id uint
			if idFloat, ok := idVal.(float64); ok {
				id = uint(idFloat)
				fmt.Printf("ID convertido com sucesso: %d\n", id)
			} else {
				fmt.Printf("ERRO: Campo 'id' não é um número válido: %T = %v\n", idVal, idVal)
				fmt.Printf("========================\n")
				return nil
			}

			// Reconstrói o objeto de usuário a partir das claims
			userDTO := &dtos.UserResponseDTO{
				UserID:    id,
				UserEmail: emailVal.(string),
			}

			fmt.Printf("Usuário reconstruído: ID=%d, Email=%s\n", userDTO.UserID, userDTO.UserEmail)
			fmt.Printf("========================\n")

			return userDTO
		},

		// Authorizator: define regras de autorização após a autenticação
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			isAdmin, exists := claims["is_admin"]
			if exists && isAdmin == true {
				return true
			}
			// Aqui você pode adicionar lógica extra para não-admins depois
			return false
		},

		// LoginResponse: personaliza a resposta HTTP em caso de login bem-sucedido
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			fmt.Println("==== Login Bem-Sucedido ====")
			fmt.Printf("Token gerado: %s\n", token)
			fmt.Printf("Expira em: %v\n", expire)

			// Recupera o objeto user do contexto (armazenado pelo Authenticator)
			user, _ := c.Get("user")

			c.JSON(code, gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
				"user":   user,
			})
		},

		// Unauthorized: personaliza a resposta HTTP em caso de falha na autenticação
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Printf("==== Falha na Autenticação ====\n")
			fmt.Printf("Rota: %s %s\n", c.Request.Method, c.Request.URL.Path)
			fmt.Printf("Código: %d, Mensagem: %s\n", code, message)

			// Retorna um erro JSON com o código e mensagem apropriados
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
}
