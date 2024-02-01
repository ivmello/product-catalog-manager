package rabbitmq

import (
	"fmt"

	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/message_broker"
)

type service struct{}

func NewRabbitMQ(db dependency_provider.DependencyProvider) message_broker.MessageBroker {
	return service{}
}

func (s service) Consumer(destinations []string, msgChan chan string) error {
	fmt.Println("consumer")
	return nil
}

func (s service) Producer(msg string, destinations []string) error {
	fmt.Println("producer")
	return nil
}
