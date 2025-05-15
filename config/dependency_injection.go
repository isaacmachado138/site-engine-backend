package config

import (
	"myapp/application/services"
	"myapp/infrastructure/repositories"

	"gorm.io/gorm"
)

type AppDependencies struct {
	UserService             *services.UserService
	SiteService             *services.SiteService
	ModuleService           *services.ModuleService
	ComponentService        *services.ComponentService
	ComponentSettingService *services.ComponentSettingService
}

func SetupDependencies(db *gorm.DB) *AppDependencies {
	userRepo := repositories.NewUserRepository(db)
	siteRepo := repositories.NewSiteRepository(db)
	moduleRepo := repositories.NewModuleRepository(db)
	componentRepo := repositories.NewComponentRepository(db)
	componentSettingRepo := repositories.NewComponentSettingRepository(db)

	return &AppDependencies{
		UserService:             services.NewUserService(userRepo),
		SiteService:             services.NewSiteService(siteRepo, moduleRepo, componentRepo),
		ModuleService:           services.NewModuleService(moduleRepo),
		ComponentService:        services.NewComponentService(componentRepo),
		ComponentSettingService: services.NewComponentSettingService(componentSettingRepo),
	}
}
