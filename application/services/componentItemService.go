package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

type ComponentItemService struct {
	repo repositories.ComponentItemRepository
}

func NewComponentItemService(repo repositories.ComponentItemRepository) *ComponentItemService {
	return &ComponentItemService{repo: repo}
}

func (s *ComponentItemService) UpsertMany(dto dtos.ComponentItemUpsertManyDTO) error {
	var items []entities.ComponentItem
	for _, i := range dto.Items {
		items = append(items, entities.ComponentItem{
			ComponentID:               i.ComponentID,
			ComponentItemTitle:        i.ComponentItemTitle,
			ComponentItemSubtitle:     i.ComponentItemSubtitle,
			ComponentItemSubtitleType: i.ComponentItemSubtitleType,
			ComponentItemText:         i.ComponentItemText,
			ComponentItemImage:        i.ComponentItemImage,
			ComponentItemOrder:        i.ComponentItemOrder,
		})
	}
	return s.repo.UpsertMany(dto.ComponentID, items)
}

func (s *ComponentItemService) FindByComponentID(componentID uint) ([]dtos.ComponentItemDTO, error) {
	items, err := s.repo.FindByComponentID(componentID)
	if err != nil {
		return nil, err
	}
	var dtosItems []dtos.ComponentItemDTO
	for _, i := range items {
		dtosItems = append(dtosItems, dtos.ComponentItemDTO{
			ComponentItemID:           i.ComponentItemID,
			ComponentID:               i.ComponentID,
			ComponentItemTitle:        i.ComponentItemTitle,
			ComponentItemSubtitle:     i.ComponentItemSubtitle,
			ComponentItemSubtitleType: i.ComponentItemSubtitleType,
			ComponentItemText:         i.ComponentItemText,
			ComponentItemImage:        i.ComponentItemImage,
			ComponentItemOrder:        i.ComponentItemOrder,
			ComponentItemLink:         i.ComponentItemLink,
		})
	}
	return dtosItems, nil
}
