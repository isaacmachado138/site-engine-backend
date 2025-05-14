package entities

// SiteModuleComponent representa a associação entre módulos e componentes
type ModuleComponent struct {
	ID          uint      `gorm:"primaryKey;column:module_component_id"`
	ModuleID    uint      `gorm:"column:module_id;not null"`
	ComponentID uint      `gorm:"column:component_id;not null"`
	Order       int       `gorm:"column:module_component_order;default:0"`
	Active      bool      `gorm:"column:module_component_active;default:true"`
	Module      Module    `gorm:"foreignKey:ModuleID;references:module_id"`
	Component   Component `gorm:"foreignKey:ComponentID;references:component_id"`
}

func (ModuleComponent) TableName() string {
	return "module_component"
}
