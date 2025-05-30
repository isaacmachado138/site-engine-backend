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
	CategoryService         *services.CategoryService
	SiteCategoryService     *services.SiteCategoryService
	CityService             *services.CityService
	JWTMiddleware           *jwt.GinJWTMiddleware
}

func SetupDependencies(db *gorm.DB, cfg *Config) *AppDependencies {
	userRepo := repositories.NewUserRepository(db)
	siteRepo := repositories.NewSiteRepository(db)
	moduleRepo := repositories.NewModuleRepository(db)
	componentRepo := repositories.NewComponentRepository(db)
	componentSettingRepo := repositories.NewComponentSettingRepository(db)
	componentItemRepo := repositories.NewComponentItemRepository(db)
	componentTypeSettingRepo := repositories.NewComponentTypeSettingRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	siteCategoryRepo := repositories.NewSiteCategoryRepository(db)
	cityRepo := repositories.NewCityRepository(db)

	userService := services.NewUserService(userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	siteCategoryService := services.NewSiteCategoryService(siteCategoryRepo)
	cityService := services.NewCityService(cityRepo)
	jwtMiddleware, _ := middlewares.SetupJWTMiddleware(userService, cfg.JWTSecret)

	return &AppDependencies{
		UserService:             userService,
		SiteService:             services.NewSiteService(siteRepo, moduleRepo, componentRepo),
		ModuleService:           services.NewModuleService(moduleRepo, siteRepo),
		ComponentService:        services.NewComponentService(componentRepo, componentTypeSettingRepo, componentSettingRepo, siteRepo),
		ComponentSettingService: services.NewComponentSettingService(componentSettingRepo),
		ComponentItemService:    services.NewComponentItemService(componentItemRepo),
		CategoryService:         categoryService,
		SiteCategoryService:     siteCategoryService,
		CityService:             cityService,
		JWTMiddleware:           jwtMiddleware,
	}
}
