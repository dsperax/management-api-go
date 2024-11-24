package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// ActiveProducer é responsável por publicar eventos de ativação de usuário no Kafka.
type ActiveProducer struct {
	Writer *kafka.Writer
}

// NewActiveProducer cria uma nova instância de ActiveProducer com as configurações fornecidas.
func NewActiveProducer(brokers []string, topic string) *ActiveProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &ActiveProducer{
		Writer: writer,
	}
}

// PublishUserActive publica um evento de usuário ativo no Kafka.
func (kp *ActiveProducer) PublishUserActive(name string) error {
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("User-%s", name)),
		Value: []byte(name),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := kp.Writer.WriteMessages(ctx, msg)
	if err != nil {
		return fmt.Errorf("erro ao enviar mensagem para o Kafka: %v", err)
	}
	return nil
}

// Close encerra a conexão do produtor Kafka.
func (kp *ActiveProducer) Close() error {
	return kp.Writer.Close()
}
