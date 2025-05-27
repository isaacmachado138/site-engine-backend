package repositories

import (
	"myapp/domain/entities"
)

type ICategoryRepository interface {
	Create(category *entities.Category) error
	Update(category *entities.Category) error
	Delete(id uint) error
	FindByID(id uint) (*entities.Category, error)
	FindAll() ([]entities.Category, error)
	FindActive() ([]entities.Category, error)
}
