package dtos

// ComponentCreateDTO representa os dados necessários para criar um componente
type ComponentCreateDTO struct {
	ComponentTypeId string `json:"component_type_id" binding:"required"`
	ComponentName   string `json:"component_name" binding:"required"`
	UserID          uint   `json:"user_id" binding:"required"`
}

// ComponentResponseDTO representa os dados retornados após a criação de um componente
type ComponentResponseDTO struct {
	ComponentID       uint                   `json:"component_id"`
	ComponentTypeId   string                 `json:"component_type_id"`
	ComponentTypeCode string                 `json:"component_type_code"`
	ComponentName     string                 `json:"component_name"`
	UserId            uint                   `json:"user_id"`
	ComponentSettings map[string]interface{} `json:"component_settings"`
}
