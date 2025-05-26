package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComponentItemHandler struct {
	service          *services.ComponentItemService
	componentService *services.ComponentService
}

func NewComponentItemHandler(service *services.ComponentItemService, componentService *services.ComponentService) *ComponentItemHandler {
	return &ComponentItemHandler{
		service:          service,
		componentService: componentService,
	}
}

func (h *ComponentItemHandler) UpsertMany(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Verificar se o componente pertence ao usuário logado
	if err := h.componentService.VerifyOwnership(uint(componentID), uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Verificar se o componente pertence ao usuário logado
	if err := h.componentService.VerifyOwnership(uint(componentID), uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := h.service.FindByComponentID(uint(componentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
