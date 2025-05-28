package repositories

import "myapp/domain/entities"

// SiteFilters define os filtros disponíveis para busca de sites
type SiteFilters struct {
	UserID     *string `json:"user_id,omitempty"`
	CategoryID *uint   `json:"category_id,omitempty"`
	Active     *bool   `json:"active,omitempty"`
	// Futuros filtros podem ser adicionados aqui
	// LocationID *uint   `json:"location_id,omitempty"`
	// SearchTerm *string `json:"search_term,omitempty"`
}

// SiteRepository define os métodos para o repositório de sites
type SiteRepository interface {
	Create(site *entities.Site) error
	FindBySlug(slug string) (*entities.Site, error)
	FindByUserID(userId string) ([]entities.Site, error)
	FindWithFilters(filters SiteFilters) ([]entities.Site, error) // Novo método genérico
	FindByID(siteID uint) (*entities.Site, error)
	Update(site *entities.Site) error
}
