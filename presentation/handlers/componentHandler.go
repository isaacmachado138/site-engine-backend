package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myapp/application/dtos"
	"myapp/application/services"
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
