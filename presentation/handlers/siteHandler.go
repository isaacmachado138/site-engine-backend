package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
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

// GetSites - Método genérico para buscar sites com diferentes filtros dinâmicos
// Suporta parâmetros: ?user_id=123, ?category_id=456, ?active=true, etc.
func (h *SiteHandler) GetSites(c *gin.Context) {
	// Extrair parâmetros de query
	userID := c.Query("user_id")
	categoryIDStr := c.Query("category_id")
	activeStr := c.Query("active")

	// Validar se pelo menos um parâmetro foi fornecido
	/*if userID == "" && categoryIDStr == "" && activeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Pelo menos um parâmetro de busca é obrigatório (user_id, category_id, active, etc.)",
		})
		return
	}*/

	// Construir filtros dinamicamente
	var filters repositories.SiteFilters

	// Aplicar filtro de user_id se fornecido
	if userID != "" {
		filters.UserID = &userID
	}

	// Aplicar filtro de category_id se fornecido
	if categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category_id deve ser um número válido"})
			return
		}
		categoryIDUint := uint(categoryID)
		filters.CategoryID = &categoryIDUint
	}

	// Aplicar filtro de active se fornecido
	if activeStr != "" {
		active := activeStr == "true" || activeStr == "1"
		filters.Active = &active
	}

	// Chamar serviço genérico
	sites, err := h.siteService.GetSitesWithFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sites)
}

// Update lida com a atualização parcial de sites
func (h *SiteHandler) Update(c *gin.Context) {
	// Extrair user ID do contexto
	userIDStr, exists := c.Get("user_id_logged")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	var updateDTO dtos.SiteUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	siteIDParam := c.Param("siteId")
	var siteID uint
	_, err = fmt.Sscanf(siteIDParam, "%d", &siteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Verificar se o site pertence ao usuário logado
	if err := h.siteService.VerifyOwnership(siteID, uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	site, err := h.siteService.Update(siteID, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, site)
}
