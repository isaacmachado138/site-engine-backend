package entities

// ComponentTypeSetting representa a associação entre tipos de componentes e suas configurações disponíveis
type ComponentTypeSetting struct {
	ComponentTypeID uint   `gorm:"column:component_type_id;primaryKey;not null"`
	SettingKey      string `gorm:"column:component_setting_key;primaryKey;size:100;not null"`

	// Relacionamentos
	ComponentType *ComponentType `gorm:"foreignKey:ComponentTypeID;references:ID"`
}

func (ComponentTypeSetting) TableName() string {
	return "component_type_setting"
}
