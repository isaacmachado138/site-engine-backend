package entities

// User representa um usuário do sistema
type User struct {
	ID       uint   `gorm:"column:user_id;primaryKey"`
	Name     string `gorm:"column:user_name;size:100;not null"`
	Email    string `gorm:"column:user_email;size:100;not null;unique"`
	Password string `gorm:"column:user_password;size:255;not null"`
	Del      int    `gorm:"column:user_del;default:0"`
	Sites    []Site `gorm:"foreignKey:UserID;references:user_id"`
}

func (User) TableName() string {
	return "user"
}
