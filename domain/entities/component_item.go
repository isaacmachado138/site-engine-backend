package entities

type ComponentItem struct {
	ComponentItemID           uint   `gorm:"column:component_item_id;primaryKey" json:"component_item_id"`
	ComponentID               uint   `gorm:"column:component_id" json:"component_id"`
	ComponentItemTitle        string `gorm:"column:component_item_title" json:"component_item_title"`
	ComponentItemSubtitle     string `gorm:"column:component_item_subtitle" json:"component_item_subtitle"`
	ComponentItemSubtitleType string `gorm:"column:component_item_subtitle_type" json:"component_item_subtitle_type"`
	ComponentItemText         string `gorm:"column:component_item_text" json:"component_item_text"`
	ComponentItemImage        string `gorm:"column:component_item_image" json:"component_item_image"`
	ComponentItemOrder        int    `gorm:"column:component_item_order" json:"component_item_order"`
	ComponentItemLink         string `gorm:"column:component_item_link" json:"component_item_link"`
}

func (ComponentItem) TableName() string {
	return "component_item"
}
