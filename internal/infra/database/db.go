package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

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

func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}
}
