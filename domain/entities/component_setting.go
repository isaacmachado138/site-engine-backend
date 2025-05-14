package entities

// ComponentSiteSetting representa configurações específicas de componentes em sites
type ComponentSetting struct {
	ID          uint   `gorm:"primaryKey;column:component_setting_id"`
	ComponentID uint   `gorm:"column:component_id;not null"`
	Key         string `gorm:"column:component_setting_key;size:100;not null"`
	Value       string `gorm:"column:component_setting_value;type:text"`
}

func (ComponentSetting) TableName() string {
	return "component_setting"
}
