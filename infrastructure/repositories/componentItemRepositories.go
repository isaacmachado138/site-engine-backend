package repositories

import (
	"myapp/domain/entities"

	"gorm.io/gorm"
)

type componentItemRepository struct {
	db *gorm.DB
}

func NewComponentItemRepository(db *gorm.DB) *componentItemRepository {
	return &componentItemRepository{db: db}
}

func (r *componentItemRepository) UpsertMany(componentID uint, items []entities.ComponentItem) error {
	// Remove todos os itens antigos e insere os novos (simples, pode ser otimizado)
	if err := r.db.Where("component_id = ?", componentID).Delete(&entities.ComponentItem{}).Error; err != nil {
		return err
	}
	if len(items) > 0 {
		return r.db.Create(&items).Error
	}
	return nil
}

func (r *componentItemRepository) FindByComponentID(componentID uint) ([]entities.ComponentItem, error) {
	var items []entities.ComponentItem
	err := r.db.Where("component_id = ?", componentID).Order("component_item_order asc").Find(&items).Error
	return items, err
}
