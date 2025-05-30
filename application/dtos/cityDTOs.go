package dtos

// CityResponseDTO representa os dados de resposta de uma cidade
type CityResponseDTO struct {
	CityID   uint   `json:"city_id"`
	CityName string `json:"city_name"`
}
