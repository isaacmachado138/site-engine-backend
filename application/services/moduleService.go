package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

// ModuleService lida com operações relacionadas a módulos
type ModuleService struct {
	moduleRepository repositories.ModuleRepository
}

// NewModuleService cria uma nova instância de ModuleService
func NewModuleService(moduleRepository repositories.ModuleRepository) *ModuleService {
	return &ModuleService{
		moduleRepository: moduleRepository,
	}
}

// Create cria um novo módulo
func (s *ModuleService) Create(moduleDTO dtos.ModuleCreateDTO) (*dtos.ModuleResponseDTO, error) {
	module := &entities.Module{
		Name:        moduleDTO.ModuleName,
		Slug:        moduleDTO.ModuleSlug,
		Description: moduleDTO.ModuleDescription,
		Order:       moduleDTO.ModuleOrder,
		SiteID:      moduleDTO.SiteID,
	}

	if err := s.moduleRepository.Create(module); err != nil {
		return nil, err
	}

	return &dtos.ModuleResponseDTO{
		ModuleID:          module.ID,
		ModuleName:        module.Name,
		ModuleSlug:        module.Slug,
		ModuleDescription: module.Description,
		ModuleOrder:       module.Order,
		SiteID:            module.SiteID,
	}, nil
}
