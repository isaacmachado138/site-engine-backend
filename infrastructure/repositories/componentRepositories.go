package repositories

import (
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"

	"gorm.io/gorm"
)

// componentRepository implementa a interface ComponentRepository
type componentRepository struct {
	db *gorm.DB
}

// NewComponentRepository cria uma nova instância do repositório de componentes
func NewComponentRepository(db *gorm.DB) repositories.ComponentRepository {
	return &componentRepository{
		db: db,
	}
}

// Create insere um novo componente no banco de dados
func (r *componentRepository) Create(component *entities.Component) error {
	return r.db.Create(component).Error
}

// FindByModuleID busca todos os componentes de um módulo via tabela module_component
func (r *componentRepository) FindByModuleID(moduleID uint) ([]entities.Component, error) {
	// Buscar os relacionamentos module_component com o ID do módulo especificado
	var moduleComponents []entities.ModuleComponent
	if err := r.db.Where("module_id = ?", moduleID).Find(&moduleComponents).Error; err != nil {
		return nil, err
	}

	// Verificar se há componentes relacionados
	if len(moduleComponents) == 0 {
		return []entities.Component{}, nil
	}

	// Extrair os IDs dos componentes
	var componentIDs []uint
	for _, mc := range moduleComponents {
		componentIDs = append(componentIDs, mc.ComponentID)
	}

	// Buscar todos os componentes em uma única consulta, incluindo o tipo e os itens
	var components []entities.Component
	if err := r.db.Preload("Settings").Preload("Type").Preload("Items").Where("component_id IN ?", componentIDs).Find(&components).Error; err != nil {
		return nil, err
	}

	return components, nil
}

// Busca componente único por site e code exato
func (r *componentRepository) FindUniqueBySiteAndTypeCode(siteID uint, typeCode string) (*entities.Component, error) {
	var component entities.Component
	err := r.db.Preload("Settings").Preload("Type").Joins("JOIN component_type ON component.component_type_id = component_type.component_type_id").
		Where("component.site_id = ? AND component_type.component_type_code = ? AND component_type.component_type_unique_in_site = 1", siteID, typeCode).
		First(&component).Error
	if err != nil {
		return nil, err
	}
	return &component, nil
}

// Busca componente único por site e code LIKE (ex: navbar, footer)
func (r *componentRepository) FindUniqueBySiteAndTypeCodeLike(siteID uint, typeCodeLike string) (*entities.Component, error) {
	var component entities.Component
	err := r.db.Preload("Settings").Preload("Type").Joins("JOIN component_type ON component.component_type_id = component_type.component_type_id").
		Where("component.site_id = ? AND component_type.component_type_code LIKE ? AND component_type.component_type_unique_in_site = 1", siteID, "%"+typeCodeLike+"%").
		First(&component).Error
	if err != nil {
		return nil, err
	}
	return &component, nil
}

// FindByID busca um componente pelo ID
func (r *componentRepository) FindByID(componentID uint) (*entities.Component, error) {
	var component entities.Component
	if err := r.db.Preload("Type").Where("component_id = ?", componentID).First(&component).Error; err != nil {
		return nil, err
	}
	return &component, nil
}
