package main

import (
	"product-catalog-manager/internal/application/configuration"
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/infra/akafka"
	"product-catalog-manager/internal/infra/http"

	"github.com/IBM/sarama"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		panic(err)
	}
	dp := dependency_provider.New(config)
	msgChan := make(chan *sarama.ConsumerMessage)
	go akafka.Consumer([]string{config.KafkaTopics}, config.KafkaServers, msgChan)
	go http.InitializeServer(dp)
	for msg := range msgChan {
		dp.GetProductService().HandleMessage(string(msg.Value))
	}
}
