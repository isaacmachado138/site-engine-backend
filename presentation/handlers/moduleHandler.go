package handlers

import (
	"fmt"
	"net/http"

	"myapp/application/dtos"
	"myapp/application/services"

	"github.com/gin-gonic/gin"
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

// Update lida com a atualização de módulos existentes
func (h *ModuleHandler) Update(c *gin.Context) {
	var updateDTO dtos.ModuleUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	moduleIDParam := c.Param("moduleId")
	var moduleID uint
	_, err := fmt.Sscanf(moduleIDParam, "%d", &moduleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	module, err := h.moduleService.Update(moduleID, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}
