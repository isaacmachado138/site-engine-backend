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

// GetCategoriesBySite busca todas as categorias associadas a um site
func (s *SiteCategoryService) GetCategoriesBySite(siteID uint) ([]uint, error) {
	siteCategories, err := s.repo.FindBySiteID(siteID)
	if err != nil {
		return nil, err
	}

	categoryIDs := make([]uint, len(siteCategories))
	for i, sc := range siteCategories {
		categoryIDs[i] = sc.CategoryID
	}

	return categoryIDs, nil
}

// UpdateSiteCategories atualiza todas as categorias de um site
func (s *SiteCategoryService) UpdateSiteCategories(siteID uint, categoryIDs []uint) error {
	// Primeiro, remove todas as categorias existentes do site
	if err := s.repo.DeleteBySiteID(siteID); err != nil {
		return err
	}

	// Depois, adiciona as novas categorias
	for _, categoryID := range categoryIDs {
		siteCategory := &entities.SiteCategory{
			SiteID:     siteID,
			CategoryID: categoryID,
		}
		if err := s.repo.Create(siteCategory); err != nil {
			return err
		}
	}

	return nil
}
