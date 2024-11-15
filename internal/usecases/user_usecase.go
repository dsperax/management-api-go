package usecases

import (
	"github.com/dsperax/management-api-go/internal/domain/entities"
	"github.com/dsperax/management-api-go/internal/domain/repositories"
)

type KafkaProducer interface {
	PublishUserActive(name string) error
}

type UserUseCase struct {
	UserRepo      repositories.UserRepository
	KafkaProducer KafkaProducer
}

func (u *UserUseCase) GetUser(id int64) (*entities.User, error) {
	user, err := u.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Publica o nome do usuário no tópico Kafka
	err = u.KafkaProducer.PublishUserActive(user.Name)
	if err != nil {
		// Trate o erro conforme necessário (ex: log)
		// Não interrompemos a execução caso o envio falhe
	}

	return user, nil
}
