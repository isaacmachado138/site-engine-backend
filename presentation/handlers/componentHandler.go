package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"myapp/application/dtos"
	"myapp/application/services"

	"github.com/gin-gonic/gin"
)

// ComponentHandler lida com operações relacionadas a componentes
type ComponentHandler struct {
	componentService *services.ComponentService
}

// NewComponentHandler cria uma nova instância de ComponentHandler
func NewComponentHandler(componentService *services.ComponentService) *ComponentHandler {
	return &ComponentHandler{
		componentService: componentService,
	}
}

// Create lida com a criação de novos componentes
func (h *ComponentHandler) Create(c *gin.Context) {
	var componentDTO dtos.ComponentCreateDTO
	if err := c.ShouldBindJSON(&componentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	component, err := h.componentService.Create(componentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, component)
}

// GetByID busca um componente específico pelo ID
func (h *ComponentHandler) GetByID(c *gin.Context) {
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

	componentID := c.Param("componentId")
	if componentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do componente é obrigatório"})
		return
	}

	// Converter string para uint
	var id uint
	if _, err := fmt.Sscanf(componentID, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do componente inválido"})
		return
	}

	// Verificar se o componente pertence ao usuário logado
	if err := h.componentService.VerifyOwnership(id, uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	component, err := h.componentService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Componente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, component)
}
