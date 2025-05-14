package entities

// Component representa um componente genérico
type Component struct {
	ID       uint               `gorm:"column:component_id;primaryKey"`
	Type     string             `gorm:"column:component_type;size:100;not null"`
	Name     string             `gorm:"column:component_name;size:100;not null"`
	UserID   uint               `gorm:"column:user_id;not null"`
	ModuleID uint               `gorm:"column:component_module_id;not null"`
	Module   Module             `gorm:"foreignKey:ModuleID;references:ID"`
	Settings []ComponentSetting `gorm:"foreignKey:ComponentID;references:ID"`
}

func (Component) TableName() string {
	return "component"
}
