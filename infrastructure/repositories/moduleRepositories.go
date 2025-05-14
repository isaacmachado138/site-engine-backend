package repositories

import (
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"

	"gorm.io/gorm"
)

// moduleRepository implementa a interface ModuleRepository
type moduleRepository struct {
	db *gorm.DB
}

// NewModuleRepository cria uma nova instância do repositório de módulos
func NewModuleRepository(db *gorm.DB) repositories.ModuleRepository {
	return &moduleRepository{
		db: db,
	}
}

// Create insere um novo módulo no banco de dados
func (r *moduleRepository) Create(module *entities.Module) error {
	return r.db.Create(module).Error
}

func (r *moduleRepository) FindBySiteID(siteID uint) ([]entities.Module, error) {
	var modules []entities.Module
	if err := r.db.Where("site_id = ?", siteID).Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}
