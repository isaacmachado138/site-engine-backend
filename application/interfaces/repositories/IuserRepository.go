package repositories

import (
	"myapp/domain/entities"
)

// UserRepository define as operações possíveis no repositório de usuários
type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id uint) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	List() ([]*entities.User, error)
	IsFirstUser() (bool, error)
}
