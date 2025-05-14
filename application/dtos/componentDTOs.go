package dtos

// ComponentCreateDTO representa os dados necessários para criar um componente
type ComponentCreateDTO struct {
	ComponentType string `json:"component_type" binding:"required"`
	ComponentName string `json:"component_name" binding:"required"`
	UserID        uint   `json:"user_id" binding:"required"`
}

// ComponentResponseDTO representa os dados retornados após a criação de um componente
type ComponentResponseDTO struct {
	ComponentID   uint   `json:"component_id"`
	ComponentType string `json:"component_type"`
	ComponentName string `json:"component_name"`
	UserId        uint   `json:"user_id"`
}
