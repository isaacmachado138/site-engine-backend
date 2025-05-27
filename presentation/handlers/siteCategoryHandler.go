package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SiteCategoryHandler struct {
	siteCategoryService *services.SiteCategoryService
}

func NewSiteCategoryHandler(siteCategoryService *services.SiteCategoryService) *SiteCategoryHandler {
	return &SiteCategoryHandler{siteCategoryService: siteCategoryService}
}

// Associa uma categoria a um site
func (h *SiteCategoryHandler) AddCategoryToSite(c *gin.Context) {
	var dto dtos.SiteCategoryCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := h.siteCategoryService.AddCategoryToSite(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Associação criada"})
}

// Remove uma categoria de um site
func (h *SiteCategoryHandler) RemoveCategoryFromSite(c *gin.Context) {
	var dto dtos.SiteCategoryCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := h.siteCategoryService.RemoveCategoryFromSite(dto.SiteID, dto.CategoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Associação removida"})
}
