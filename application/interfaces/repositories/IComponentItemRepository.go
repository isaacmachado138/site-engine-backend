package repositories

import "myapp/domain/entities"

type ComponentItemRepository interface {
	UpsertMany(componentID uint, items []entities.ComponentItem) error
	FindByComponentID(componentID uint) ([]entities.ComponentItem, error)
}
