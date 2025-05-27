package entities

// Category representa a tabela de categorias
type Category struct {
	ID          uint   `gorm:"column:category_id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:category_name;size:100;not null"`
	Description string `gorm:"column:category_description;type:text"`
	Active      bool   `gorm:"column:category_active;default:1"`
	Icon        string `gorm:"column:category_icon;type:text"`
	Main        bool   `gorm:"column:category_main;default:0"`
	Sites       []Site `gorm:"many2many:site_category;joinForeignKey:CategoryID;JoinReferences:SiteID"`
}

func (Category) TableName() string {
	return "category"
}
