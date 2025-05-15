package repositories

import (
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"

	"gorm.io/gorm"
)

type componentSettingRepository struct {
	db *gorm.DB
}

func NewComponentSettingRepository(db *gorm.DB) repositories.ComponentSettingRepository {
	return &componentSettingRepository{db: db}
}

func (r *componentSettingRepository) CreateMany(componentID uint, settings []entities.ComponentSetting) error {
	for i := range settings {
		settings[i].ComponentID = componentID
	}
	return r.db.Create(&settings).Error
}

func (r *componentSettingRepository) UpdateMany(componentID uint, settings []entities.ComponentSetting) error {
	for _, s := range settings {
		r.db.Model(&entities.ComponentSetting{}).
			Where("component_id = ? AND component_setting_key = ?", componentID, s.Key).
			Update("component_setting_value", s.Value)
	}
	return nil
}

func (r *componentSettingRepository) DeleteByComponentID(componentID uint) error {
	return r.db.Where("component_id = ?", componentID).Delete(&entities.ComponentSetting{}).Error
}

func (r *componentSettingRepository) FindByComponentID(componentID uint) ([]entities.ComponentSetting, error) {
	var settings []entities.ComponentSetting
	err := r.db.Where("component_id = ?", componentID).Find(&settings).Error
	return settings, err
}
