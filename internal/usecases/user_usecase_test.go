package usecases

import (
	"fmt"
	"testing"

	"github.com/dsperax/management-api-go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório de usuários
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(id int64) (*entities.User, error) {
	args := m.Called(id)
	userInterface := args.Get(0)
	if userInterface != nil {
		user := userInterface.(*entities.User)
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock do KafkaProducer
type MockKafkaProducer struct {
	mock.Mock
}

func (m *MockKafkaProducer) PublishUserActive(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func TestGetUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockProducer := new(MockKafkaProducer)
	user := &entities.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Configura o comportamento esperado dos mocks
	mockRepo.On("FindByID", int64(1)).Return(user, nil)
	mockProducer.On("PublishUserActive", user.Name).Return(nil)

	// Cria o use case com os mocks injetados
	useCase := UserUseCase{
		UserRepo:      mockRepo,
		KafkaProducer: mockProducer,
	}

	// Act
	result, err := useCase.GetUser(1)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}

func TestGetUser_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockProducer := new(MockKafkaProducer)

	// Configura o comportamento esperado dos mocks
	mockRepo.On("FindByID", int64(1)).Return(nil, fmt.Errorf("usuário não encontrado"))

	// Cria o use case com os mocks injetados
	useCase := UserUseCase{
		UserRepo:      mockRepo,
		KafkaProducer: mockProducer,
	}

	// Act
	result, err := useCase.GetUser(1)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "usuário não encontrado")
	mockRepo.AssertExpectations(t)
	// Não esperamos chamadas ao KafkaProducer neste caso
	mockProducer.AssertNotCalled(t, "PublishUserActive")
}

func TestGetUser_KafkaPublishError(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockProducer := new(MockKafkaProducer)
	user := &entities.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Configura o comportamento esperado dos mocks
	mockRepo.On("FindByID", int64(1)).Return(user, nil)
	mockProducer.On("PublishUserActive", user.Name).Return(fmt.Errorf("erro ao publicar no Kafka"))

	// Cria o use case com os mocks injetados
	useCase := UserUseCase{
		UserRepo:      mockRepo,
		KafkaProducer: mockProducer,
	}

	// Act
	result, err := useCase.GetUser(1)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}

func TestGetUser_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockProducer := new(MockKafkaProducer)

	// Configura o comportamento esperado dos mocks
	mockRepo.On("FindByID", int64(1)).Return(nil, fmt.Errorf("erro no repositório"))

	// Cria o use case com os mocks injetados
	useCase := UserUseCase{
		UserRepo:      mockRepo,
		KafkaProducer: mockProducer,
	}

	// Act
	result, err := useCase.GetUser(1)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "erro no repositório")
	mockRepo.AssertExpectations(t)
	// Não esperamos chamadas ao KafkaProducer neste caso
	mockProducer.AssertNotCalled(t, "PublishUserActive")
}
