package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

type ComponentSettingService struct {
	repo repositories.ComponentSettingRepository
}

func NewComponentSettingService(repo repositories.ComponentSettingRepository) *ComponentSettingService {
	return &ComponentSettingService{repo: repo}
}

// Cria ou atualiza vários settings para um componente
func (s *ComponentSettingService) UpsertMany(componentID uint, dtosSettings []dtos.ComponentSettingCreateDTO) error {
	var settings []entities.ComponentSetting
	for _, dto := range dtosSettings {
		settings = append(settings, entities.ComponentSetting{
			ComponentID: componentID,
			Key:         dto.ComponentSettingKey,
			Value:       dto.ComponentSettingValue,
		})
	}
	// Agora faz upsert real: insere se não existir, atualiza se existir
	return s.repo.UpsertMany(componentID, settings)
}

func (s *ComponentSettingService) GetByComponentID(componentID uint) ([]dtos.ComponentSettingResponseDTO, error) {
	settings, err := s.repo.FindByComponentID(componentID)
	if err != nil {
		return nil, err
	}
	var resp []dtos.ComponentSettingResponseDTO
	for _, s := range settings {
		resp = append(resp, dtos.ComponentSettingResponseDTO{
			ID:                    s.ID,
			ComponentID:           s.ComponentID,
			ComponentSettingKey:   s.Key,
			ComponentSettingValue: s.Value,
		})
	}
	return resp, nil
}
