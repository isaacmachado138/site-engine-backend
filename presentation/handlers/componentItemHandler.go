package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComponentItemHandler struct {
	service *services.ComponentItemService
}

func NewComponentItemHandler(service *services.ComponentItemService) *ComponentItemHandler {
	return &ComponentItemHandler{service: service}
}

func (h *ComponentItemHandler) UpsertMany(c *gin.Context) {
	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var items []dtos.ComponentItemDTO
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	// Validação: todos os component_id devem ser iguais ao da URL
	for _, item := range items {
		if item.ComponentID != uint(componentID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os component_id devem ser iguais ao da URL"})
			return
		}
	}
	dto := dtos.ComponentItemUpsertManyDTO{
		ComponentID: uint(componentID),
		Items:       items,
	}
	if err := h.service.UpsertMany(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ComponentItemHandler) FindByComponentID(c *gin.Context) {
	componentIDStr := c.Param("componentId")
	componentID, err := strconv.ParseUint(componentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	items, err := h.service.FindByComponentID(uint(componentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
