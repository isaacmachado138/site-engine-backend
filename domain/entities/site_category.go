package entities

// SiteCategory representa a tabela de ligação site_category
type SiteCategory struct {
	SiteID     uint `gorm:"column:site_id;primaryKey"`
	CategoryID uint `gorm:"column:category_id;primaryKey"`
}

func (SiteCategory) TableName() string {
	return "site_category"
}
