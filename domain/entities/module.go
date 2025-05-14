package entities

// Module representa um módulo associado a um site
type Module struct {
	ID          uint        `gorm:"primaryKey;column:module_id"`
	Name        string      `gorm:"column:module_name;size:100;not null"`
	Slug        string      `gorm:"column:module_slug;size:100;not null;unique"`
	Description string      `gorm:"column:module_description;type:text"`
	Order       int         `gorm:"column:module_order;default:0"`
	SiteID      uint        `gorm:"column:site_id;not null"`
	Site        Site        `gorm:"foreignKey:SiteID;references:site_id"`
	Components  []Component `gorm:"foreignKey:ModuleID;references:module_id"`
}

func (Module) TableName() string {
	return "module"
}
