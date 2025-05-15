package entities

// Component representa um componente genérico
type Component struct {
	ID       uint               `gorm:"column:component_id;primaryKey"`
	TypeId   string             `gorm:"column:component_type_id;size:100;not null"`
	Type     *ComponentType     `gorm:"foreignKey:TypeId;references:ID"`
	Name     string             `gorm:"column:component_name;size:100;not null"`
	SiteId   uint               `gorm:"column:site_id;not null"`
	UserID   uint               `gorm:"column:user_id;not null"`
	Settings []ComponentSetting `gorm:"foreignKey:ComponentID;references:ID"`
}

func (Component) TableName() string {
	return "component"
}
