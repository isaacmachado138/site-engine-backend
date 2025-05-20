package dtos

type ComponentItemDTO struct {
	ComponentItemID           uint   `json:"component_item_id"`
	ComponentID               uint   `json:"component_id" binding:"required"`
	ComponentItemTitle        string `json:"component_item_title" binding:"required"`
	ComponentItemSubtitle     string `json:"component_item_subtitle"`
	ComponentItemSubtitleType string `json:"component_item_subtitle_type"`
	ComponentItemText         string `json:"component_item_text"`
	ComponentItemImage        string `json:"component_item_image"`
	ComponentItemOrder        int    `json:"component_item_order"`
}

type ComponentItemUpsertManyDTO struct {
	ComponentID uint               `json:"component_id" binding:"required"`
	Items       []ComponentItemDTO `json:"items" binding:"required"`
}
