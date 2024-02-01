package rabbitmq_adapter

import (
	"fmt"
	"log"

	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/message_broker"

	"github.com/wagslane/go-rabbitmq"
)

type service struct {
	dp dependency_provider.DependencyProvider
}

func NewRabbitMQ(db dependency_provider.DependencyProvider) message_broker.MessageBroker {
	return service{
		dp: db,
	}
}

func (s service) Consumer(destinations []string, msgChan chan string) error {
	fmt.Println("consumer")
	return nil
}

func (s service) Producer(msg string, destinations []string) error {
	fmt.Println("producer")
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@rabbitmq:5672",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("product_catalog_manager"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	err = publisher.Publish(
		[]byte("hello, world"),
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
	if err != nil {
		log.Println(err)
	}
	return nil
}
