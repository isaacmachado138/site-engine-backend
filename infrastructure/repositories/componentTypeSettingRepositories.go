package repositories

import (
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"

	"gorm.io/gorm"
)

type componentTypeSettingRepository struct {
	db *gorm.DB
}

func NewComponentTypeSettingRepository(db *gorm.DB) repositories.ComponentTypeSettingRepository {
	return &componentTypeSettingRepository{db: db}
}

// FindByComponentTypeID busca todas as configurações disponíveis para um tipo de componente
func (r *componentTypeSettingRepository) FindByComponentTypeID(componentTypeID uint) ([]entities.ComponentTypeSetting, error) {
	var settings []entities.ComponentTypeSetting
	err := r.db.Where("component_type_id = ?", componentTypeID).Find(&settings).Error
	return settings, err
}

// FindByComponentTypeCode busca todas as configurações disponíveis para um tipo de componente pelo código
func (r *componentTypeSettingRepository) FindByComponentTypeCode(componentTypeCode string) ([]entities.ComponentTypeSetting, error) {
	var settings []entities.ComponentTypeSetting
	err := r.db.Joins("JOIN component_type ON component_type.component_type_id = component_type_setting.component_type_id").
		Where("component_type.component_type_code = ?", componentTypeCode).
		Find(&settings).Error
	return settings, err
}

// GetAllWithSettings busca todos os tipos de componentes com suas configurações disponíveis
func (r *componentTypeSettingRepository) GetAllWithSettings() ([]entities.ComponentType, error) {
	var componentTypes []entities.ComponentType

	// Usando subquery para carregar as configurações relacionadas
	err := r.db.Preload("Settings", func(db *gorm.DB) *gorm.DB {
		return db.Table("component_type_setting").Select("component_type_id, component_setting_key")
	}).Find(&componentTypes).Error

	return componentTypes, err
}

// ValidateSettingForType verifica se uma configuração é válida para um tipo de componente
func (r *componentTypeSettingRepository) ValidateSettingForType(componentTypeID uint, settingKey string) (bool, error) {
	var count int64
	err := r.db.Model(&entities.ComponentTypeSetting{}).
		Where("component_type_id = ? AND component_setting_key = ?", componentTypeID, settingKey).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetSettingKeysByTypeCode busca apenas as chaves de configuração para um tipo específico
func (r *componentTypeSettingRepository) GetSettingKeysByTypeCode(componentTypeCode string) ([]string, error) {
	var settingKeys []string

	err := r.db.Table("component_type_setting").
		Select("component_type_setting.component_setting_key").
		Joins("JOIN component_type ON component_type.component_type_id = component_type_setting.component_type_id").
		Where("component_type.component_type_code = ?", componentTypeCode).
		Pluck("component_setting_key", &settingKeys).Error

	return settingKeys, err
}
