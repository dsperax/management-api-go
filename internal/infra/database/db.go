package database

import (
	"database/sql"
	"fmt"

	// Import necessário para registrar o driver MySQL
	_ "github.com/go-sql-driver/mysql"
)

// DB é a instância global da conexão com o banco de dados.
var DB *sql.DB

// InitDB inicializa a conexão com o banco de dados.
func InitDB() error {
	dsn := "user:userpassword@tcp(localhost:3306)/management_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao abrir o banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	DB = db
	return nil
}

// CloseDB fecha a conexão com o banco de dados.
func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}
}
