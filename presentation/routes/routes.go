package routes

import (
	"myapp/application/services"
	"myapp/presentation/handlers"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configura todas as rotas da API
func SetupRoutes(router *gin.Engine, userService *services.UserService, siteService *services.SiteService, moduleService *services.ModuleService, componentService *services.ComponentService) {
	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userService)
	siteHandler := handlers.NewSiteHandler(siteService)
	moduleHandler := handlers.NewModuleHandler(moduleService)
	componentHandler := handlers.NewComponentHandler(componentService)

	// Definir grupo base da API
	api := router.Group("/api")
	{
		// Rotas de usuários
		api.POST("/user/register", userHandler.Register)

		// Rotas de sites
		api.POST("/site", siteHandler.Create)

		// Rotas de módulos
		api.POST("/site/module", moduleHandler.Create)

		// Rotas de componentes
		api.POST("/site/component", componentHandler.Create)

		// Rota para buscar site por slug
		api.GET("/site/:slug", siteHandler.GetBySlug)
	}
}

// setupHealthRoutes configura rotas de health check
func setupHealthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/health", func(context *gin.Context) {
		// Verificação do banco de dados
		sqlDB, err := db.DB()

		// Informações do banco
		dbInfo := gin.H{
			"status": "desconectado",
		}

		if err != nil {
			dbInfo["error"] = err.Error()
		} else {
			// Testar conexão com ping
			pingErr := sqlDB.Ping()

			// Coletar estatísticas
			stats := sqlDB.Stats()

			dbInfo = gin.H{
				"status": func() string {
					if pingErr == nil {
						return "conectado"
					} else {
						return "erro"
					}
				}(),
				"conexoes_abertas":  stats.OpenConnections,
				"conexoes_em_uso":   stats.InUse,
				"conexoes_idle":     stats.Idle,
				"tempo_espera_max":  stats.MaxIdleClosed,
				"tempo_espera":      stats.WaitDuration.String(),
				"conexoes_max_vida": stats.MaxLifetimeClosed,
			}

			if pingErr != nil {
				dbInfo["erro_ping"] = pingErr.Error()
			}

			// Consultar versão do banco (opcional)
			var version string
			row := db.Raw("SELECT version()").Row()
			if row.Scan(&version) == nil {
				dbInfo["version"] = version
			}
		}

		// Informações da aplicação
		appInfo := gin.H{
			"status":    "operacional",
			"timestamp": time.Now().Format(time.RFC3339),
		}

		context.JSON(http.StatusOK, gin.H{
			"aplicacao":   appInfo,
			"banco_dados": dbInfo,
		})
	})
}

// setupAuthRoutes configura rotas de autenticação
func setupAuthRoutes(router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	auth := router.Group("/auth")
	{
		// Rotas públicas de autenticação
		auth.POST("/login", authMiddleware.LoginHandler)
		auth.GET("/refresh", authMiddleware.RefreshHandler)
	}
}
