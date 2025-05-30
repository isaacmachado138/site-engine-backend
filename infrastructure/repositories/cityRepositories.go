package repositories

import (
	"myapp/domain/entities"

	"gorm.io/gorm"
)

// CityRepository implementa a interface ICityRepository
type CityRepository struct {
	db *gorm.DB
}

// NewCityRepository cria uma nova instância de CityRepository
func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}

// FindAll busca todas as cidades
func (r *CityRepository) FindAll() ([]entities.City, error) {
	var cities []entities.City
	if err := r.db.Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
