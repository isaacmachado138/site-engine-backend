package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

	var updateDTO dtos.ModuleUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	moduleIDParam := c.Param("moduleId")
	var moduleID uint
	_, err = fmt.Sscanf(moduleIDParam, "%d", &moduleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Verificar se o módulo pertence ao usuário logado
	if err := h.moduleService.VerifyOwnership(moduleID, uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module, err := h.moduleService.Update(moduleID, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}
