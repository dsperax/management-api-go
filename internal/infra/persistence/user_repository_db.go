package persistence

import (
	"database/sql"
	"fmt"

	"github.com/dsperax/management-api-go/internal/domain/entities"
	"github.com/dsperax/management-api-go/internal/domain/repositories"
	"github.com/dsperax/management-api-go/internal/infra/database"
)

// UserRepositoryDB implementa UserRepository usando um banco de dados relacional.
type UserRepositoryDB struct {
	db *sql.DB
}

// NewUserRepositoryDB cria uma nova instância de UserRepositoryDB com o banco de dados fornecido.
func NewUserRepositoryDB() repositories.UserRepository {
	return &UserRepositoryDB{
		db: database.DB,
	}
}

// FindByID busca um usuário pelo ID no banco de dados.
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
