package repositories

import "myapp/domain/entities"

// ComponentTypeSettingRepository define os métodos para o repositório de configurações de tipos de componentes
type ComponentTypeSettingRepository interface {
	// FindByComponentTypeID busca todas as configurações disponíveis para um tipo de componente
	FindByComponentTypeID(componentTypeID uint) ([]entities.ComponentTypeSetting, error)

	// FindByComponentTypeCode busca todas as configurações disponíveis para um tipo de componente pelo código
	FindByComponentTypeCode(componentTypeCode string) ([]entities.ComponentTypeSetting, error)

	// GetAllWithSettings busca todos os tipos de componentes com suas configurações disponíveis
	GetAllWithSettings() ([]entities.ComponentType, error)

	// ValidateSettingForType verifica se uma configuração é válida para um tipo de componente
	ValidateSettingForType(componentTypeID uint, settingKey string) (bool, error)

	// GetSettingKeysByTypeCode busca apenas as chaves de configuração para um tipo específico
	GetSettingKeysByTypeCode(componentTypeCode string) ([]string, error)
}
