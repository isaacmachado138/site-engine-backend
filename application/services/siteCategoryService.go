package services

import (
	"errors"
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

type SiteCategoryService struct {
	repo repositories.ISiteCategoryRepository
}

func NewSiteCategoryService(repo repositories.ISiteCategoryRepository) *SiteCategoryService {
	return &SiteCategoryService{repo: repo}
}

func (s *SiteCategoryService) AddCategoryToSite(dto dtos.SiteCategoryCreateDTO) error {
	existing, err := s.repo.FindBySiteAndCategory(dto.SiteID, dto.CategoryID)
	if err == nil && existing != nil {
		return errors.New("Associação já existe")
	}
	return s.repo.Create(&entities.SiteCategory{
		SiteID:     dto.SiteID,
		CategoryID: dto.CategoryID,
	})
}

func (s *SiteCategoryService) RemoveCategoryFromSite(siteID, categoryID uint) error {
	return s.repo.Delete(siteID, categoryID)
}
