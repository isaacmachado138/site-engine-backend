package repositories

import "myapp/domain/entities"

// ComponentSettingRepository define os métodos para o repositório de settings de componentes
type ComponentSettingRepository interface {
	CreateMany(componentID uint, settings []entities.ComponentSetting) error
	UpdateMany(componentID uint, settings []entities.ComponentSetting) error
	DeleteByComponentID(componentID uint) error
	FindByComponentID(componentID uint) ([]entities.ComponentSetting, error)
}
