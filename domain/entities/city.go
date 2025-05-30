package entities

// City representa uma cidade no sistema
type City struct {
	CityID   uint   `gorm:"primaryKey;column:city_id" json:"city_id"`
	CityName string `gorm:"column:city_name;type:varchar(100);not null" json:"city_name"`
}

// TableName especifica o nome da tabela no banco de dados
func (City) TableName() string {
	return "city"
}
