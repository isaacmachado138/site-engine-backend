package services

import (
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
)

// CityService lida com operações relacionadas a cidades
type CityService struct {
	cityRepository repositories.ICityRepository
}

// NewCityService cria uma nova instância de CityService
func NewCityService(cityRepository repositories.ICityRepository) *CityService {
	return &CityService{
		cityRepository: cityRepository,
	}
}

// GetAll busca todas as cidades
func (s *CityService) GetAll() ([]dtos.CityResponseDTO, error) {
	cities, err := s.cityRepository.FindAll()
	if err != nil {
		return []dtos.CityResponseDTO{}, err
	}

	// Inicializar como array vazio para garantir que retorna [] ao invés de null
	resp := make([]dtos.CityResponseDTO, 0)

	for _, city := range cities {
		resp = append(resp, dtos.CityResponseDTO{
			CityID:   city.CityID,
			CityName: city.CityName,
		})
	}

	return resp, nil
}
