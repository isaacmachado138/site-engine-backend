package handlers

import (
	"myapp/application/dtos"
	"myapp/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// GetActiveCategories retorna todas as categorias ativas
func (h *CategoryHandler) GetActiveCategories(c *gin.Context) {
	categories, err := h.categoryService.FindActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// Create insere uma nova categoria
func (h *CategoryHandler) Create(c *gin.Context) {
	var dto dtos.CategoryCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	cat, err := h.categoryService.Create(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}
