package dtos

// ComponentTypeSettingDTO representa os dados de configuração disponível para um tipo de componente
type ComponentTypeSettingDTO struct {
	ComponentTypeID uint   `json:"component_type_id"`
	SettingKey      string `json:"component_setting_key"`
}

// ComponentTypeWithSettingsDTO representa um tipo de componente com suas configurações disponíveis
type ComponentTypeWithSettingsDTO struct {
	ComponentTypeID           uint     `json:"component_type_id"`
	ComponentTypeCode         string   `json:"component_type_code"`
	ComponentTypeDescription  string   `json:"component_type_description,omitempty"`
	ComponentTypeUniqueInSite bool     `json:"component_type_unique_in_site,omitempty"`
	AvailableSettings         []string `json:"available_settings"`
}

// ComponentTypeSettingResponseDTO representa a resposta de configurações por tipo
type ComponentTypeSettingResponseDTO struct {
	ComponentTypeID   uint     `json:"component_type_id"`
	ComponentTypeCode string   `json:"component_type_code"`
	SettingKeys       []string `json:"setting_keys"`
}
