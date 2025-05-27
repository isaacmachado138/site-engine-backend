package routes

import (
	"myapp/config"
	"myapp/presentation/handlers"
	"myapp/presentation/middlewares"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configura todas as rotas da API
func SetupRoutes(router *gin.Engine, deps interface{}) {
	// Type assertion para o struct de dependências
	appDeps := deps.(*config.AppDependencies)
	userHandler := handlers.NewUserHandler(appDeps.UserService)
	siteHandler := handlers.NewSiteHandler(appDeps.SiteService)
	moduleHandler := handlers.NewModuleHandler(appDeps.ModuleService)
	componentHandler := handlers.NewComponentHandler(appDeps.ComponentService)
	componentSettingHandler := handlers.NewComponentSettingHandler(appDeps.ComponentSettingService, appDeps.ComponentService)
	componentItemHandler := handlers.NewComponentItemHandler(appDeps.ComponentItemService, appDeps.ComponentService)
	categoryHandler := handlers.NewCategoryHandler(appDeps.CategoryService) // Novo manipulador para categorias
	// Usa setupAuthRoutes para configurar rotas de autenticação (apenas login e refresh agora)
	setupAuthRoutes(router, appDeps.JWTMiddleware) // Definir grupo base da API protegido
	api := router.Group("/api")
	api.Use(middlewares.TokenExtractor()) // Adiciona o TokenExtractor corretamente
	api.Use(appDeps.JWTMiddleware.MiddlewareFunc())
	api.Use(middlewares.TokenDebugger())   // Adiciona debug das claims do token
	api.Use(middlewares.UserIDExtractor()) // Extrai o user_id_logged das claims
	{
		// Rota de criação de usuário agora exige autenticação e admin
		api.POST("/user/register", middlewares.AdminRequired(), userHandler.Register)

		api.POST("/categories", categoryHandler.Create)

		// Rotas de sites
		api.POST("/site", siteHandler.Create)
		api.PUT("/site/:siteId", siteHandler.Update)

		// Rotas de módulos
		api.POST("/site/module", moduleHandler.Create)
		api.PUT("/site/module/:moduleId", moduleHandler.Update)
		// Rotas de componentes
		api.POST("/site/component", componentHandler.Create)
		api.GET("/site/component/:componentId", componentHandler.GetByID)

		// Rota para settings de componentes
		api.POST("/site/component/:componentId/setting", componentSettingHandler.UpsertMany)
		api.GET("/site/component/:componentId/setting", componentSettingHandler.FindByComponentID)

		// Rotas para items de componentes
		api.POST("/site/component/:componentId/items", componentItemHandler.UpsertMany)
		api.GET("/site/component/:componentId/items", componentItemHandler.FindByComponentID)
	}

	// Rota pública para buscar site por slug
	router.GET("/api/site/:slug", siteHandler.GetBySlug)

	// Rota para buscar todos os sites de um usuário
	router.GET("/api/:userId/sites", siteHandler.GetSitesByUser)

	// Rotas para categorias
	router.GET("/categories/active", categoryHandler.GetActiveCategories)
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

// setupAuthRoutes configura rotas de autenticação (login e refresh)
func setupAuthRoutes(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	auth := router.Group("/api/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh", authMiddleware.RefreshHandler)
}
