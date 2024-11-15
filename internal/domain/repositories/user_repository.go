package repositories

import "github.com/dsperax/management-api-go/internal/domain/entities"

type UserRepository interface {
	FindByID(id int64) (*entities.User, error)
	// Outros métodos como Save, Update, Delete
}
