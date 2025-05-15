package dtos

// ComponentSettingCreateDTO representa o payload para criar/atualizar settings de um componente
// Usado para requests em lote
// Exemplo de uso: [{component_setting_key: "key", component_setting_value: "value"}, ...]
type ComponentSettingCreateDTO struct {
	ComponentSettingKey   string `json:"component_setting_key" binding:"required"`
	ComponentSettingValue string `json:"component_setting_value" binding:"required"`
}

type ComponentSettingResponseDTO struct {
	ID                    uint   `json:"id"`
	ComponentID           uint   `json:"component_id"`
	ComponentSettingKey   string `json:"component_setting_key"`
	ComponentSettingValue string `json:"component_setting_value"`
}
