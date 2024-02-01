package main

import (
	"product-catalog-manager/internal/configuration"
	"product-catalog-manager/internal/dependency_provider"
	"product-catalog-manager/internal/infra/akafka"
	"product-catalog-manager/internal/infra/rest_api"

	"github.com/IBM/sarama"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		panic(err)
	}
	dp := dependency_provider.New(config)
	msgChan := make(chan *sarama.ConsumerMessage)
	go akafka.Consumer([]string{"product"}, config.KafkaServers, msgChan)
	go rest_api.InitializeServer(dp)
	for msg := range msgChan {
		dp.GetProductService().HandleMessage(string(msg.Value))
	}
}
