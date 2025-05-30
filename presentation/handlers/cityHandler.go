package handlers

import (
	"myapp/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	cityService *services.CityService
}

func NewCityHandler(cityService *services.CityService) *CityHandler {
	return &CityHandler{cityService: cityService}
}

// GetAll busca todas as cidades
func (h *CityHandler) GetAll(c *gin.Context) {
	cities, err := h.cityService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cities)
}
