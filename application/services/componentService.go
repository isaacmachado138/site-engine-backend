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
		TypeId: componentDTO.ComponentTypeId,
		Name:   componentDTO.ComponentName,
		UserID: componentDTO.UserID,
	}

	if err := s.componentRepository.Create(component); err != nil {
		return nil, err
	}

	// Buscar o componente com as informações do tipo
	createdComponent, err := s.componentRepository.FindByID(component.ID)
	if err != nil {
		return nil, err
	}

	// Obter o código do tipo de componente
	typeCode := ""
	if createdComponent.Type != nil {
		typeCode = createdComponent.Type.Code
	}

	return &dtos.ComponentResponseDTO{
		ComponentID:       createdComponent.ID,
		ComponentTypeId:   createdComponent.TypeId,
		ComponentTypeCode: typeCode,
		ComponentName:     createdComponent.Name,
		UserId:            createdComponent.UserID,
	}, nil
}
