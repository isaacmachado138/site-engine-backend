package repositories

import (
	"myapp/domain/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}

func (r *CategoryRepository) FindByID(id uint) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) FindActive() ([]entities.Category, error) {
	var categories []entities.Category
	if err := r.db.Where("category_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
