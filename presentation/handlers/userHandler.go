package handlers

import (
	"net/http"

	"myapp/application/dtos"
	"myapp/application/services"

	"github.com/gin-gonic/gin"
)

// UserHandler lida com operações relacionadas a usuários
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler cria uma nova instância de UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register lida com o registro de novos usuários
func (h *UserHandler) Register(c *gin.Context) {
	var userDTO dtos.UserCreateDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	user, err := h.userService.Create(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
