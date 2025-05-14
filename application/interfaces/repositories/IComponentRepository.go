package repositories

import "myapp/domain/entities"

// ComponentRepository define os métodos para o repositório de componentes
type ComponentRepository interface {
	Create(component *entities.Component) error
	FindByModuleID(moduleID uint) ([]entities.Component, error)
}
