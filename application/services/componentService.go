package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

// ComponentService lida com operações relacionadas a componentes
type ComponentService struct {
	componentRepository      repositories.ComponentRepository
	componentTypeSettingRepo repositories.ComponentTypeSettingRepository
	componentSettingRepo     repositories.ComponentSettingRepository
}

// NewComponentService cria uma nova instância de ComponentService
func NewComponentService(
	componentRepository repositories.ComponentRepository,
	componentTypeSettingRepo repositories.ComponentTypeSettingRepository,
	componentSettingRepo repositories.ComponentSettingRepository,
) *ComponentService {
	return &ComponentService{
		componentRepository:      componentRepository,
		componentTypeSettingRepo: componentTypeSettingRepo,
		componentSettingRepo:     componentSettingRepo,
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
	var typeIdUint uint
	if createdComponent.Type != nil {
		typeCode = createdComponent.Type.Code
		typeIdUint = createdComponent.Type.ID
	}

	// Buscar as setting_keys disponíveis para esse tipo de componente
	settings, err := s.componentTypeSettingRepo.FindByComponentTypeID(typeIdUint)
	if err != nil {
		return nil, err
	}

	// Criar settings vazios para o novo componente
	var settingsToInsert []entities.ComponentSetting
	for _, sKey := range settings {
		settingsToInsert = append(settingsToInsert, entities.ComponentSetting{
			ComponentID: createdComponent.ID,
			Key:         sKey.SettingKey,
			Value:       "",
		})
	}
	if len(settingsToInsert) > 0 {
		err = s.componentSettingRepo.CreateMany(createdComponent.ID, settingsToInsert)
		if err != nil {
			return nil, err
		}
	}

	return &dtos.ComponentResponseDTO{
		ComponentID:       createdComponent.ID,
		ComponentTypeId:   createdComponent.TypeId,
		ComponentTypeCode: typeCode,
		ComponentName:     createdComponent.Name,
		UserId:            createdComponent.UserID,
	}, nil
}
