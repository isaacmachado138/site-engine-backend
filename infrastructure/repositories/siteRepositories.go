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
	if err := r.db.Where("site_slug = ?", slug).First(&site).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) FindByUserID(userId string) ([]entities.Site, error) {
	var sites []entities.Site
	err := r.db.Where("user_id = ?", userId).Find(&sites).Error
	return sites, err
}

// FindByID busca um site pelo ID
func (r *siteRepository) FindByID(siteID uint) (*entities.Site, error) {
	var site entities.Site
	if err := r.db.First(&site, siteID).Error; err != nil {
		return nil, err
	}
	return &site, nil
}

// Update atualiza um site existente
func (r *siteRepository) Update(site *entities.Site) error {
	return r.db.Save(site).Error
}
