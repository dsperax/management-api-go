package main

import (
	"log"

	"github.com/dsperax/management-api-go/internal/infra/database"
	"github.com/dsperax/management-api-go/internal/interfaces/routes"

	_ "github.com/dsperax/management-api-go/docs" // Importa os documentos gerados pelo Swagger
	"github.com/gin-gonic/gin"
)

// @title           Management API
// @version         1.0
// @description     API de gerenciamento de usuários.

// @host      localhost:8080
// @BasePath  /

// @schemes http
func main() {
	// Inicializa a conexão com o banco de dados
	if err := database.InitDB(); err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}
	defer database.CloseDB()

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
