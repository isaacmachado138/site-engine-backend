package handlers

import (
	"fmt"
	"net/http"

	"myapp/application/dtos"
	"myapp/application/services"

	jwt "github.com/appleboy/gin-jwt/v2"
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
	claims := jwt.ExtractClaims(c)
	fmt.Printf("[DEBUG] Claims recebidas no handler Register: %+v\n", claims)
	if len(claims) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou ausente"})
		return
	}

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
