package repositories

import "myapp/domain/entities"

// ICityRepository define a interface para operações de cidade
type ICityRepository interface {
	FindAll() ([]entities.City, error)
}
