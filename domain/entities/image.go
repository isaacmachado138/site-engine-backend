package entities

// Image representa uma imagem vinculada a um usuário
type Image struct {
	ID     uint   `gorm:"primaryKey;column:image_id"`
	URL    string `gorm:"column:image_url;size:255;not null"`
	Alt    string `gorm:"column:image_alt;size:255"`
	Title  string `gorm:"column:image_title;size:255"`
	UserID uint   `gorm:"column:user_id;not null"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (Image) TableName() string {
	return "image"
}
