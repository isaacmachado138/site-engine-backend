package dtos

// ModuleCreateDTO representa os dados necessários para criar um módulo
type ModuleCreateDTO struct {
	ModuleName        string `json:"module_name" binding:"required"`
	ModuleSlug        string `json:"module_slug" binding:"required"`
	ModuleDescription string `json:"module_description"`
	ModuleOrder       int    `json:"module_order"`
	SiteID            uint   `json:"site_id" binding:"required"`
}

// ModuleResponseDTO representa os dados retornados após a criação de um módulo
type ModuleResponseDTO struct {
	ModuleID          uint   `json:"module_id"`
	ModuleName        string `json:"module_name"`
	ModuleSlug        string `json:"module_slug"`
	ModuleDescription string `json:"module_description"`
	ModuleOrder       int    `json:"module_order"`
	SiteID            uint   `json:"site_id"`
}
