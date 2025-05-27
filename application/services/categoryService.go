package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

type CategoryService struct {
	repo repositories.ICategoryRepository
}

func NewCategoryService(repo repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(dto dtos.CategoryCreateDTO) (*dtos.CategoryResponseDTO, error) {
	category := &entities.Category{
		Name:        dto.Name,
		Description: dto.Description,
		Active:      dto.Active,
		Icon:        dto.Icon,
		Main:        dto.Main,
	}
	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return &dtos.CategoryResponseDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Active:      category.Active,
		Icon:        category.Icon,
		Main:        category.Main,
	}, nil
}

func (s *CategoryService) Update(id uint, dto dtos.CategoryUpdateDTO) (*dtos.CategoryResponseDTO, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, nil
	}

	if dto.Name != nil {
		category.Name = *dto.Name
	}
	if dto.Description != nil {
		category.Description = *dto.Description
	}
	if dto.Active != nil {
		category.Active = *dto.Active
	}
	if dto.Icon != nil {
		category.Icon = *dto.Icon
	}
	if dto.Main != nil {
		category.Main = *dto.Main
	}

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}
	return &dtos.CategoryResponseDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Active:      category.Active,
		Icon:        category.Icon,
		Main:        category.Main,
	}, nil
}

func (s *CategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) FindByID(id uint) (*dtos.CategoryResponseDTO, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, nil
	}
	return &dtos.CategoryResponseDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Active:      category.Active,
		Icon:        category.Icon,
		Main:        category.Main,
	}, nil
}

func (s *CategoryService) FindAll() ([]dtos.CategoryResponseDTO, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.CategoryResponseDTO
	for _, category := range categories {
		result = append(result, dtos.CategoryResponseDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Active:      category.Active,
			Icon:        category.Icon,
			Main:        category.Main,
		})
	}
	return result, nil
}

func (s *CategoryService) FindActive() ([]dtos.CategoryResponseDTO, error) {
	categories, err := s.repo.FindActive()
	if err != nil {
		return nil, err
	}
	var result []dtos.CategoryResponseDTO
	for _, category := range categories {
		result = append(result, dtos.CategoryResponseDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Active:      category.Active,
			Icon:        category.Icon,
			Main:        category.Main,
		})
	}
	return result, nil
}
