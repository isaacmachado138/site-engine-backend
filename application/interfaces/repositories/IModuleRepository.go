package repositories

import "myapp/domain/entities"

// ModuleRepository define os métodos para o repositório de módulos
type ModuleRepository interface {
	Create(module *entities.Module) error
	FindBySiteID(siteID uint) ([]entities.Module, error)
	FindByID(moduleID uint) (*entities.Module, error)
	Update(module *entities.Module) error
}
