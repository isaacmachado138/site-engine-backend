package dtos

// SiteCreateDTO representa os dados necessários para criar um site
type SiteCreateDTO struct {
	SiteName       string `json:"site_name" binding:"required"`
	SiteSlug       string `json:"site_slug" binding:"required"`
	UserID         uint   `json:"user_id" binding:"required"`
	SiteIconWindow string `json:"site_icon_window"`
}

// SiteResponseDTO representa os dados retornados após a criação de um site
type SiteResponseDTO struct {
	SiteID         uint   `json:"site_id"`
	SiteName       string `json:"site_name"`
	SiteSlug       string `json:"site_slug"`
	SiteIconWindow string `json:"site_icon_window"`
}

// SiteFullResponseDTO representa o site completo com módulos e componentes
type SiteFullResponseDTO struct {
	SiteID         uint                      `json:"site_id"`
	SiteName       string                    `json:"site_name"`
	SiteSlug       string                    `json:"site_slug"`
	SiteIconWindow string                    `json:"site_icon_window"`
	Modules        []ModuleWithComponentsDTO `json:"modules"`
	Navbar         *ComponentDTO             `json:"navbar,omitempty"`
	Footer         *ComponentDTO             `json:"footer,omitempty"`
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

// ComponentTypeDTO representa os dados do tipo de componente
// para resposta em endpoints que retornam componentes
// Pode ser expandido conforme necessário
//
type ComponentTypeDTO struct {
	ComponentTypeId           uint   `json:"component_type_id"`
	ComponentTypeCode         string `json:"component_type_code"`
	ComponentTypeDescription  string `json:"component_type_description,omitempty"`
	ComponentTypeUniqueInSite bool   `json:"component_type_unique_in_site,omitempty"`
}

// Atualize o ComponentDTO para incluir o tipo de componente
type ComponentDTO struct {
	ComponentID       uint               `json:"component_id"`
	ComponentTypeId   string             `json:"component_type_id"`
	ComponentTypeCode string             `json:"component_type_code"`
	ComponentName     string             `json:"component_name"`
	ComponentSettings map[string]string  `json:"component_settings"`
	Items             []ComponentItemDTO `json:"items"` // <-- Adicionado para trazer os items do componente
}
