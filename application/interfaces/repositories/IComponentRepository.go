package repositories

import "myapp/domain/entities"

// ComponentRepository define os métodos para o repositório de componentes
type ComponentRepository interface {
	Create(component *entities.Component) error
	FindByModuleID(moduleID uint) ([]entities.Component, error)
	FindUniqueBySiteAndTypeCode(siteID uint, typeCode string) (*entities.Component, error)
	FindUniqueBySiteAndTypeCodeLike(siteID uint, typeCodeLike string) (*entities.Component, error)
	FindByID(componentID uint) (*entities.Component, error)
}
