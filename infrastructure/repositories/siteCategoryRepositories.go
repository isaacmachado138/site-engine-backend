package repositories

import (
	"myapp/domain/entities"

	"gorm.io/gorm"
)

type SiteCategoryRepository struct {
	db *gorm.DB
}

func NewSiteCategoryRepository(db *gorm.DB) *SiteCategoryRepository {
	return &SiteCategoryRepository{db: db}
}

func (r *SiteCategoryRepository) Create(siteCategory *entities.SiteCategory) error {
	return r.db.Create(siteCategory).Error
}

func (r *SiteCategoryRepository) Delete(siteID, categoryID uint) error {
	return r.db.Delete(&entities.SiteCategory{}, "site_id = ? AND category_id = ?", siteID, categoryID).Error
}

func (r *SiteCategoryRepository) FindBySiteAndCategory(siteID, categoryID uint) (*entities.SiteCategory, error) {
	var sc entities.SiteCategory
	if err := r.db.Where("site_id = ? AND category_id = ?", siteID, categoryID).First(&sc).Error; err != nil {
		return nil, err
	}
	return &sc, nil
}
