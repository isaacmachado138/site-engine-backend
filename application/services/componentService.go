package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

// ComponentService lida com operações relacionadas a componentes
type ComponentService struct {
	componentRepository repositories.ComponentRepository
}

// NewComponentService cria uma nova instância de ComponentService
func NewComponentService(componentRepository repositories.ComponentRepository) *ComponentService {
	return &ComponentService{
		componentRepository: componentRepository,
	}
}

// Create cria um novo componente
func (s *ComponentService) Create(componentDTO dtos.ComponentCreateDTO) (*dtos.ComponentResponseDTO, error) {
	component := &entities.Component{
		Type:   componentDTO.ComponentType,
		Name:   componentDTO.ComponentName,
		UserID: componentDTO.UserID,
	}

	if err := s.componentRepository.Create(component); err != nil {
		return nil, err
	}

	return &dtos.ComponentResponseDTO{
		ComponentID:   component.ID,
		ComponentType: component.Type,
		ComponentName: component.Name,
		UserId:        component.UserID, // Corrected field name to match DTO
	}, nil
}
