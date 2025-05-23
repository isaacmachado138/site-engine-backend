package dtos

// ModuleCreateDTO representa os dados necessários para criar um módulo
// Adicionando o campo ModuleActive
type ModuleCreateDTO struct {
	ModuleName        string `json:"module_name" binding:"required"`
	ModuleSlug        string `json:"module_slug" binding:"required"`
	ModuleDescription string `json:"module_description"`
	ModuleOrder       int    `json:"module_order"`
	SiteID            uint   `json:"site_id" binding:"required"`
	ModuleActive      int    `json:"module_active"`
}

// ModuleResponseDTO representa os dados retornados após a criação de um módulo
// Adicionando o campo ModuleActive
type ModuleResponseDTO struct {
	ModuleID          uint   `json:"module_id"`
	ModuleName        string `json:"module_name"`
	ModuleSlug        string `json:"module_slug"`
	ModuleDescription string `json:"module_description"`
	ModuleOrder       int    `json:"module_order"`
	SiteID            uint   `json:"site_id"`
	ModuleActive      int    `json:"module_active"`
}

// ModuleUpdateDTO representa os dados para atualização parcial de um módulo
// Todos os campos são opcionais
// Adicione a tag 'omitempty' para permitir update parcial
type ModuleUpdateDTO struct {
	ModuleName        *string `json:"module_name,omitempty"`
	ModuleSlug        *string `json:"module_slug,omitempty"`
	ModuleDescription *string `json:"module_description,omitempty"`
	ModuleOrder       *int    `json:"module_order,omitempty"`
	SiteID            *uint   `json:"site_id,omitempty"`
	ModuleActive      *int    `json:"module_active,omitempty"`
}
