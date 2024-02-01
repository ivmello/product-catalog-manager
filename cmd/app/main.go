package main

import (
	"product-catalog-manager/internal/application/configuration"
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/infra/http_adapter"
	"product-catalog-manager/internal/infra/kafka_adapter"

	"github.com/IBM/sarama"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		panic(err)
	}
	dp := dependency_provider.New(config)
	msgChan := make(chan *sarama.ConsumerMessage)
	go kafka_adapter.Consumer([]string{config.KafkaTopics}, config.KafkaServers, msgChan)
	go http_adapter.InitializeServer(dp)
	for msg := range msgChan {
		dp.GetProductService().HandleMessage(string(msg.Value))
	}
}
