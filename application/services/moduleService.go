package services

import (
	"errors"
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
		Name:         moduleDTO.ModuleName,
		Slug:         moduleDTO.ModuleSlug,
		Description:  moduleDTO.ModuleDescription,
		Order:        moduleDTO.ModuleOrder,
		SiteID:       moduleDTO.SiteID,
		ModuleActive: moduleDTO.ModuleActive, // Propaga o campo para a entidade
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
		ModuleActive:      module.ModuleActive, // Retorna o campo no DTO
	}, nil
}

// Update atualiza um módulo existente (agora com update parcial)
func (s *ModuleService) Update(moduleID uint, updateDTO dtos.ModuleUpdateDTO) (*dtos.ModuleResponseDTO, error) {
	// Buscar o módulo atual
	module, err := s.moduleRepository.FindByID(moduleID)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, errors.New("módulo não encontrado")
	}
	// Atualizar apenas os campos enviados
	if updateDTO.ModuleName != nil {
		module.Name = *updateDTO.ModuleName
	}
	if updateDTO.ModuleSlug != nil {
		module.Slug = *updateDTO.ModuleSlug
	}
	if updateDTO.ModuleDescription != nil {
		module.Description = *updateDTO.ModuleDescription
	}
	if updateDTO.ModuleOrder != nil {
		module.Order = *updateDTO.ModuleOrder
	}
	if updateDTO.SiteID != nil {
		module.SiteID = *updateDTO.SiteID
	}
	if updateDTO.ModuleActive != nil {
		module.ModuleActive = *updateDTO.ModuleActive
	}
	if err := s.moduleRepository.Update(module); err != nil {
		return nil, err
	}
	return &dtos.ModuleResponseDTO{
		ModuleID:          module.ID,
		ModuleName:        module.Name,
		ModuleSlug:        module.Slug,
		ModuleDescription: module.Description,
		ModuleOrder:       module.Order,
		SiteID:            module.SiteID,
		ModuleActive:      module.ModuleActive,
	}, nil
}
