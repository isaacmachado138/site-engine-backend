package entities

// Site representa um site associado a um usuário
type Site struct {
	ID             uint       `gorm:"primaryKey;column:site_id"`
	Name           string     `gorm:"column:site_name;size:100;not null"`
	Slug           string     `gorm:"column:site_slug;size:100;not null;unique"`
	Description    string     `gorm:"column:site_description;type:text"`
	CityID         *uint      `gorm:"column:city_id"`
	City           *City      `gorm:"foreignKey:CityID;references:city_id"`
	HasWebsite     bool       `gorm:"column:site_has_website;default:0"`
	Keywords       string     `gorm:"column:site_keywords;type:text"`
	PhoneWhatsapp  string     `gorm:"column:site_phone_whatsapp;size:50"`
	Phone          string     `gorm:"column:site_phone;size:50"`
	UserID         uint       `gorm:"column:user_id;not null"`
	User           User       `gorm:"foreignKey:UserID;references:user_id"`
	Modules        []Module   `gorm:"foreignKey:SiteID;references:site_id"`
	SiteIconWindow string     `gorm:"column:site_icon_window;size:255"`
	Categories     []Category `gorm:"many2many:site_category;joinForeignKey:SiteID;JoinReferences:CategoryID"`
}

func (Site) TableName() string {
	return "site"
}
