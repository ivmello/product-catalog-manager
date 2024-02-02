package rabbitmq_adapter

import (
	"log"

	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/message_broker"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQConfig struct {
	ExchangeName string
	QueueName    string
	RoutingKey   string
	Message      string
}

type adapter struct {
	dp *dependency_provider.DependencyProvider
}

func NewRabbitMQAdapter(dp *dependency_provider.DependencyProvider) message_broker.MessageBroker {
	return &adapter{
		dp: dp,
	}
}

func (a *adapter) Consumer(params interface{}, msgChan chan []byte) {
	config := params.(RabbitMQConfig)
	conn, err := rabbitmq.NewConn(
		a.dp.GetConfig().RabbitMQURI,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			msgChan <- d.Body
			return rabbitmq.Ack
		},
		config.QueueName,
		rabbitmq.WithConsumerOptionsRoutingKey(config.RoutingKey),
		rabbitmq.WithConsumerOptionsExchangeName(config.ExchangeName),
		rabbitmq.WithConsumerOptionsQueueDurable,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	select {}
}

func (s *adapter) Producer(params interface{}, msgChan chan []byte) {
	return
}
