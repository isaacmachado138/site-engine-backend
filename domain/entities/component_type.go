package entities

// ComponentType representa o tipo de componente
type ComponentType struct {
	ID           uint   `gorm:"column:component_type_id;primaryKey"`
	Code         string `gorm:"column:component_type_code"`
	Description  string `gorm:"column:component_type_description"`
	UniqueInSite bool   `gorm:"column:component_type_unique_in_site"`
}

func (ComponentType) TableName() string {
	return "component_type"
}
