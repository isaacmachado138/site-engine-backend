package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myapp/application/dtos"
	"myapp/application/services"
)

// ModuleHandler lida com operações relacionadas a módulos
type ModuleHandler struct {
	moduleService *services.ModuleService
}

// NewModuleHandler cria uma nova instância de ModuleHandler
func NewModuleHandler(moduleService *services.ModuleService) *ModuleHandler {
	return &ModuleHandler{
		moduleService: moduleService,
	}
}

// Create lida com a criação de novos módulos
func (h *ModuleHandler) Create(c *gin.Context) {
	var moduleDTO dtos.ModuleCreateDTO
	if err := c.ShouldBindJSON(&moduleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	module, err := h.moduleService.Create(moduleDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, module)
}
