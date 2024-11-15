package persistence

import (
	"database/sql"
	"fmt"

	"github.com/dsperax/management-api-go/internal/domain/entities"
	"github.com/dsperax/management-api-go/internal/domain/repositories"
	"github.com/dsperax/management-api-go/internal/infra/database"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDB() repositories.UserRepository {
	return &UserRepositoryDB{
		db: database.DB,
	}
}

func (repo *UserRepositoryDB) FindByID(id int64) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var user entities.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	return &user, nil
}
