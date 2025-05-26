package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComponentSettingHandler struct {
	service          *services.ComponentSettingService
	componentService *services.ComponentService
}

func NewComponentSettingHandler(service *services.ComponentSettingService, componentService *services.ComponentService) *ComponentSettingHandler {
	return &ComponentSettingHandler{
		service:          service,
		componentService: componentService,
	}
}

// POST /api/site/component/:componentId/setting
func (h *ComponentSettingHandler) UpsertMany(c *gin.Context) {
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

	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "componentId inválido"})
		return
	}

	// Verificar se o componente pertence ao usuário logado
	if err := h.componentService.VerifyOwnership(uint(componentID), uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req []dtos.ComponentSettingCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	existing, err := h.service.GetByComponentID(uint(componentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar settings existentes"})
		return
	}

	existingKeys := make(map[string]struct{})
	for _, s := range existing {
		existingKeys[s.ComponentSettingKey] = struct{}{}
	}

	for _, novo := range req {
		if _, found := existingKeys[novo.ComponentSettingKey]; found {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe uma setting com essa key para esse componente: " + novo.ComponentSettingKey})
			return
		}
	}

	err = h.service.UpsertMany(uint(componentID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings salvos com sucesso"})
}

// GET /api/site/component/:componentId/setting
func (h *ComponentSettingHandler) FindByComponentID(c *gin.Context) {
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

	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "componentId inválido"})
		return
	}

	// Verificar se o componente pertence ao usuário logado
	if err := h.componentService.VerifyOwnership(uint(componentID), uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := h.service.GetByComponentID(uint(componentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar settings"})
		return
	}

	c.JSON(http.StatusOK, settings)
}
