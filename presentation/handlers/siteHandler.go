package handlers

import (
	"fmt"
	"net/http"

	"myapp/application/dtos"
	"myapp/application/services"

	"github.com/gin-gonic/gin"
)

// SiteHandler lida com operações relacionadas a sites
type SiteHandler struct {
	siteService *services.SiteService
}

// NewSiteHandler cria uma nova instância de SiteHandler
func NewSiteHandler(siteService *services.SiteService) *SiteHandler {
	return &SiteHandler{
		siteService: siteService,
	}
}

// Create lida com a criação de novos sites
func (h *SiteHandler) Create(c *gin.Context) {
	var siteDTO dtos.SiteCreateDTO
	if err := c.ShouldBindJSON(&siteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	site, err := h.siteService.Create(siteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, site)
}

func (h *SiteHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	onlyActive := 0
	if v := c.Query("onlyActive"); v == "1" {
		onlyActive = 1
	}

	site, err := h.siteService.GetBySlug(slug, onlyActive)
	if err != nil {
		if err.Error() == "site not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Site não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, site)
}

// GetSitesByUser retorna todos os sites de um usuário
func (h *SiteHandler) GetSitesByUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId é obrigatório"})
		return
	}
	sites, err := h.siteService.GetSitesByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sites)
}

// Update lida com a atualização parcial de sites
func (h *SiteHandler) Update(c *gin.Context) {
	var updateDTO dtos.SiteUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	siteIDParam := c.Param("siteId")
	var siteID uint
	_, err := fmt.Sscanf(siteIDParam, "%d", &siteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	site, err := h.siteService.Update(siteID, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, site)
}
