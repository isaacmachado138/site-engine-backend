package repositories

import "myapp/domain/entities"

// SiteRepository define os métodos para o repositório de sites
type SiteRepository interface {
	Create(site *entities.Site) error
	FindBySlug(slug string) (*entities.Site, error)
	FindByUserID(userId string) ([]entities.Site, error)
	FindByID(siteID uint) (*entities.Site, error)
	Update(site *entities.Site) error
}
