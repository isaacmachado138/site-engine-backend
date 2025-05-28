package handlers

import (
	"fmt"
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

// GetCategoriesBySite busca todas as categorias associadas a um site
func (h *SiteCategoryHandler) GetCategoriesBySite(c *gin.Context) {
	siteID := c.Param("siteId")
	var siteIDUint uint
	if _, err := fmt.Sscanf(siteID, "%d", &siteIDUint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do site inválido"})
		return
	}

	categoryIDs, err := h.siteCategoryService.GetCategoriesBySite(siteIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoryIDs)
}

// UpdateSiteCategories atualiza todas as categorias de um site
func (h *SiteCategoryHandler) UpdateSiteCategories(c *gin.Context) {
	siteID := c.Param("siteId")
	var siteIDUint uint
	if _, err := fmt.Sscanf(siteID, "%d", &siteIDUint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do site inválido"})
		return
	}

	var categoryIDs []uint
	if err := c.ShouldBindJSON(&categoryIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := h.siteCategoryService.UpdateSiteCategories(siteIDUint, categoryIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categorias atualizadas com sucesso"})
}
