package kafka_adapter

import (
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/message_broker"

	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Topics []string
}

type adapter struct {
	dp *dependency_provider.DependencyProvider
}

func NewKafkaAdapter(dp *dependency_provider.DependencyProvider) message_broker.MessageBroker {
	return &adapter{
		dp: dp,
	}
}

func consumePartition(consumer sarama.Consumer, topics []string, partition int32, msgChan chan string) {
	partitionConsumer, err := consumer.ConsumePartition(topics[0], partition, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	channel := partitionConsumer.Messages()
	for msg := range channel {
		msgChan <- string(msg.Value)
	}
}

func (a *adapter) Consumer(params interface{}, msgChan chan string) {
	config := sarama.NewConfig()
	consumerParams := params.(KafkaConfig)
	consumer, err := sarama.NewConsumer([]string{a.dp.GetConfig().KafkaURI}, config)
	if err != nil {
		panic(err)
	}
	partitions, err := consumer.Partitions(consumerParams.Topics[0])
	if err != nil {
		panic(err)
	}
	if len(partitions) == 0 {
		panic("No partitions found")
	}
	for _, partition := range partitions {
		consumePartition(consumer, consumerParams.Topics, partition, msgChan)
	}
}

func (a *adapter) Producer(params interface{}, msgChan chan string) {
	panic("not implemented")
}
