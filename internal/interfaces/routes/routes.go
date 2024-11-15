package routes

import (
	"github.com/dsperax/management-api-go/internal/infra/kafka"
	"github.com/dsperax/management-api-go/internal/infra/persistence"
	"github.com/dsperax/management-api-go/internal/interfaces/controllers"
	"github.com/dsperax/management-api-go/internal/usecases"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	brokers := []string{"localhost:9092"} // Use "kafka:29092" se estiver dentro do Docker
	topic := "management_user_active"
	kafkaProducer := kafka.NewActiveProducer(brokers, topic)
	userRepo := &persistence.UserRepositoryDB{}
	userUseCase := &usecases.UserUseCase{
		UserRepo:      userRepo,
		KafkaProducer: kafkaProducer,
	}
	userController := &controllers.UserController{UserUseCase: userUseCase}

	router.GET("/users/:id", userController.GetUser)
	// Outras rotas aqui

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
