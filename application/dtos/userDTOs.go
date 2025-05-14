package dtos

import (
	"myapp/domain/entities"
)

// UserCreateDTO representa os dados necessários para criar um usuário
type UserCreateDTO struct {
	UserName     string `json:"user_name" binding:"required"`
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required"`
}

// UserUpdateDTO representa os dados para atualização de um usuário
type UserUpdateDTO struct {
	UserName     string `json:"user_name" binding:"omitempty,min=3,max=100"`
	UserPassword string `json:"user_password" binding:"omitempty,min=6"`
}

// UserResponseDTO representa os dados retornados após a criação de um usuário
type UserResponseDTO struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

// ToResponseDTO converte uma entidade User para um UserResponseDTO
func ToResponseDTO(user entities.User) UserResponseDTO {
	return UserResponseDTO{
		UserID:    user.ID,
		UserName:  user.Name,
		UserEmail: user.Email,
	}
}
