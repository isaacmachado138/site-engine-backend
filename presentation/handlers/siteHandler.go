package handlers

import (
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

	// Chama o serviço para buscar o site pelo slug
	site, err := h.siteService.GetBySlug(slug)
	if err != nil {
		if err.Error() == "site not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Site não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o site encontrado
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
