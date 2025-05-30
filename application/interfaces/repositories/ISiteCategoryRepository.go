package repositories

import "myapp/domain/entities"

type ISiteCategoryRepository interface {
	Create(siteCategory *entities.SiteCategory) error
	Delete(siteID, categoryID uint) error
	FindBySiteAndCategory(siteID, categoryID uint) (*entities.SiteCategory, error)
	FindBySiteID(siteID uint) ([]entities.SiteCategory, error)
	DeleteBySiteID(siteID uint) error
}
