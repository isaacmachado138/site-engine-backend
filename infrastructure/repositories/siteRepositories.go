package repositories

import (
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"

	"gorm.io/gorm"
)

// siteRepository implementa a interface SiteRepository
type siteRepository struct {
	db *gorm.DB
}

// NewSiteRepository cria uma nova instância do repositório de sites
func NewSiteRepository(db *gorm.DB) repositories.SiteRepository {
	return &siteRepository{
		db: db,
	}
}

// Create insere um novo site no banco de dados
func (r *siteRepository) Create(site *entities.Site) error {
	return r.db.Create(site).Error
}

func (r *siteRepository) FindBySlug(slug string) (*entities.Site, error) {
	var site entities.Site
	if err := r.db.Preload("City").Where("site_slug = ?", slug).First(&site).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) FindByUserID(userId string) ([]entities.Site, error) {
	var sites []entities.Site
	err := r.db.Preload("City").Where("user_id = ?", userId).Find(&sites).Error
	return sites, err
}

// FindByID busca um site pelo ID
func (r *siteRepository) FindByID(siteID uint) (*entities.Site, error) {
	var site entities.Site
	if err := r.db.Preload("City").First(&site, siteID).Error; err != nil {
		return nil, err
	}
	return &site, nil
}

// Update atualiza um site existente
func (r *siteRepository) Update(site *entities.Site) error {
	return r.db.Save(site).Error
}

func (r *siteRepository) FindWithFilters(filters repositories.SiteFilters) ([]entities.Site, error) {
	var sites []entities.Site
	query := r.db.Model(&entities.Site{})

	// Aplicar filtro por user_id se fornecido
	if filters.UserID != nil && *filters.UserID != "" {
		query = query.Where("user_id = ?", *filters.UserID)
	}

	// Aplicar filtro por category_id se fornecido
	if filters.CategoryID != nil && *filters.CategoryID > 0 {
		query = query.Joins("JOIN site_category ON site.site_id = site_category.site_id").
			Where("site_category.category_id = ?", *filters.CategoryID)
	}

	// Aplicar filtro por nome se fornecido (busca com LIKE)
	if filters.Name != nil && *filters.Name != "" {
		query = query.Where("site_name LIKE ?", "%"+*filters.Name+"%")
	}

	// Aplicar filtro por city_id se fornecido
	if filters.CityID != nil && *filters.CityID > 0 {
		query = query.Where("city_id = ?", *filters.CityID)
	}

	// Aplicar filtro por status ativo se fornecido (para futuro uso)
	if filters.Active != nil {
		// Quando implementarmos o campo active na tabela site
		// query = query.Where("site_active = ?", *filters.Active)
	}
	// Executar a consulta
	err := query.Preload("City").Find(&sites).Error
	return sites, err
}
