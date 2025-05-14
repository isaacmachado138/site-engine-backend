package dtos

// SiteCreateDTO representa os dados necessários para criar um site
type SiteCreateDTO struct {
	SiteName string `json:"site_name" binding:"required"`
	SiteSlug string `json:"site_slug" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
}

// SiteResponseDTO representa os dados retornados após a criação de um site
type SiteResponseDTO struct {
	SiteID   uint   `json:"site_id"`
	SiteName string `json:"site_name"`
	SiteSlug string `json:"site_slug"`
}

// SiteFullResponseDTO representa o site completo com módulos e componentes
type SiteFullResponseDTO struct {
	SiteID   uint                      `json:"site_id"`
	SiteName string                    `json:"site_name"`
	SiteSlug string                    `json:"site_slug"`
	Modules  []ModuleWithComponentsDTO `json:"modules"`
}

type ModuleWithComponentsDTO struct {
	ModuleID   uint           `json:"module_id"`
	ModuleName string         `json:"module_name"`
	ModuleSlug string         `json:"module_slug"`
	Components []ComponentDTO `json:"components"`
}

type ComponentSettingDTO struct {
	ID    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ComponentDTO struct {
	ComponentID       uint              `json:"component_id"`
	ComponentType     string            `json:"component_type"`
	ComponentName     string            `json:"component_name"`
	ComponentSettings map[string]string `json:"component_settings"`
}
