package main

import (
	"product-catalog-manager/internal/application/configuration"
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/infra/http_adapter"
	"product-catalog-manager/internal/infra/kafka_adapter"
	"product-catalog-manager/internal/infra/rabbitmq_adapter"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		panic(err)
	}
	dp := dependency_provider.New(config)
	msgChan := make(chan string)
	go kafkaConsumer(dp, msgChan)
	go rabbitMQConsumer(dp, msgChan)
	go http_adapter.InitializeServer(dp)
	for msg := range msgChan {
		dp.GetProductService().HandleMessage(msg)
	}
}

func kafkaConsumer(dp *dependency_provider.DependencyProvider, msgChan chan string) {
	kafkaAdapter := kafka_adapter.NewKafkaAdapter(dp)
	kafkaAdapter.Consumer(kafka_adapter.KafkaConfig{
		Topics: []string{"products"},
	}, msgChan)
}

func rabbitMQConsumer(dp *dependency_provider.DependencyProvider, msgChan chan string) {
	rabbitmqAdapter := rabbitmq_adapter.NewRabbitMQAdapter(dp)
	rabbitmqAdapter.Consumer(rabbitmq_adapter.RabbitMQConfig{
		ExchangeName: "amq.topic",
		QueueName:    "product_catalog_manager",
		RoutingKey:   "products",
	}, msgChan)
}
