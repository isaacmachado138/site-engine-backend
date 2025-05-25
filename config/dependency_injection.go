package config

import (
	"myapp/application/services"
	"myapp/infrastructure/repositories"
	"myapp/presentation/middlewares"

	jwt "github.com/appleboy/gin-jwt/v2"
	"gorm.io/gorm"
)

type AppDependencies struct {
	UserService             *services.UserService
	SiteService             *services.SiteService
	ModuleService           *services.ModuleService
	ComponentService        *services.ComponentService
	ComponentSettingService *services.ComponentSettingService
	ComponentItemService    *services.ComponentItemService
	JWTMiddleware           *jwt.GinJWTMiddleware
}

func SetupDependencies(db *gorm.DB, cfg *Config) *AppDependencies {
	userRepo := repositories.NewUserRepository(db)
	siteRepo := repositories.NewSiteRepository(db)
	moduleRepo := repositories.NewModuleRepository(db)
	componentRepo := repositories.NewComponentRepository(db)
	componentSettingRepo := repositories.NewComponentSettingRepository(db)
	componentItemRepo := repositories.NewComponentItemRepository(db)
	//componentTypeSettingRepo := repositories.NewComponentTypeSettingRepository(db)

	userService := services.NewUserService(userRepo)
	jwtMiddleware, _ := middlewares.SetupJWTMiddleware(userService, cfg.JWTSecret)

	return &AppDependencies{
		UserService:             userService,
		SiteService:             services.NewSiteService(siteRepo, moduleRepo, componentRepo),
		ModuleService:           services.NewModuleService(moduleRepo),
		ComponentService:        services.NewComponentService(componentRepo),
		ComponentSettingService: services.NewComponentSettingService(componentSettingRepo),
		ComponentItemService:    services.NewComponentItemService(componentItemRepo),
		JWTMiddleware:           jwtMiddleware,
	}
}
