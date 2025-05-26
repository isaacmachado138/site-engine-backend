package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComponentSettingHandler struct {
	service *services.ComponentSettingService
}

func NewComponentSettingHandler(service *services.ComponentSettingService) *ComponentSettingHandler {
	return &ComponentSettingHandler{service: service}
}

// POST /api/site/component/:componentId/setting
func (h *ComponentSettingHandler) UpsertMany(c *gin.Context) {
	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "componentId inválido"})
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
	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "componentId inválido"})
		return
	}

	settings, err := h.service.GetByComponentID(uint(componentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar settings"})
		return
	}

	c.JSON(http.StatusOK, settings)
}
