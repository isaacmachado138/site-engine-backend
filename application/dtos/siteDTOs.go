package dtos

// SiteCreateDTO representa os dados necessários para criar um site
type SiteCreateDTO struct {
	SiteName          string `json:"site_name" binding:"required"`
	SiteSlug          string `json:"site_slug" binding:"required"`
	SiteDescription   string `json:"site_description"`
	CityID            *uint  `json:"city_id"`
	SiteKeywords      string `json:"site_keywords"`
	SitePhoneWhatsapp string `json:"site_phone_whatsapp"`
	SitePhone         string `json:"site_phone"`
	UserID            uint   `json:"user_id" binding:"required"`
	SiteIconWindow    string `json:"site_icon_window"`
}

// SiteResponseDTO representa os dados retornados após a criação de um site
type SiteResponseDTO struct {
	SiteID            uint   `json:"site_id"`
	SiteName          string `json:"site_name"`
	SiteSlug          string `json:"site_slug"`
	SiteDescription   string `json:"site_description"`
	CityID            *uint  `json:"city_id"`
	CityName          string `json:"city_name,omitempty"`
	SiteHasWebsite    bool   `json:"site_has_website"`
	SiteKeywords      string `json:"site_keywords"`
	SitePhoneWhatsapp string `json:"site_phone_whatsapp"`
	SitePhone         string `json:"site_phone"`
	UserID            uint   `json:"user_id"`
	SiteIconWindow    string `json:"site_icon_window"`
}

// SiteFullResponseDTO representa o site completo com módulos e componentes
type SiteFullResponseDTO struct {
	SiteID            uint                      `json:"site_id"`
	SiteName          string                    `json:"site_name"`
	SiteSlug          string                    `json:"site_slug"`
	SiteDescription   string                    `json:"site_description"`
	CityID            *uint                     `json:"city_id"`
	CityName          string                    `json:"city_name,omitempty"`
	SiteHasWebsite    bool                      `json:"site_has_website"`
	SiteKeywords      string                    `json:"site_keywords"`
	SitePhoneWhatsapp string                    `json:"site_phone_whatsapp"`
	SitePhone         string                    `json:"site_phone"`
	SiteIconWindow    string                    `json:"site_icon_window"`
	Modules           []ModuleWithComponentsDTO `json:"modules"`
	Navbar            *ComponentDTO             `json:"navbar,omitempty"`
	Footer            *ComponentDTO             `json:"footer,omitempty"`
}

type ModuleWithComponentsDTO struct {
	ModuleID          uint                   `json:"module_id"`
	ModuleName        string                 `json:"module_name"`
	ModuleSlug        string                 `json:"module_slug"`
	ModuleDescription string                 `json:"module_description"`
	ModuleOrder       int                    `json:"module_order"`
	SiteID            uint                   `json:"site_id"`
	ModuleActive      int                    `json:"module_active"`
	Components        []ComponentResponseDTO `json:"components"`
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
	ComponentID       uint                   `json:"component_id"`
	ComponentTypeId   string                 `json:"component_type_id"`
	ComponentTypeCode string                 `json:"component_type_code"`
	ComponentName     string                 `json:"component_name"`
	ComponentSettings map[string]interface{} `json:"component_settings"`
	// Items removido, pois agora vai dentro de settings
}

// SiteUpdateDTO representa os dados para atualização parcial de um site
// Todos os campos são opcionais
// Adicione a tag 'omitempty' para permitir update parcial
type SiteUpdateDTO struct {
	SiteName          *string `json:"site_name,omitempty"`
	SiteSlug          *string `json:"site_slug,omitempty"`
	SiteDescription   *string `json:"site_description,omitempty"`
	CityID            *uint   `json:"city_id,omitempty"`
	SiteKeywords      *string `json:"site_keywords,omitempty"`
	SitePhoneWhatsapp *string `json:"site_phone_whatsapp,omitempty"`
	SitePhone         *string `json:"site_phone,omitempty"`
	UserID            *uint   `json:"user_id,omitempty"`
	SiteIconWindow    *string `json:"site_icon_window,omitempty"`
}
