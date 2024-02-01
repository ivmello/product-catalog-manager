package kafka_adapter

import (
	"github.com/IBM/sarama"
)

func consumePartition(consumer sarama.Consumer, topics []string, partition int32, msgChan chan *sarama.ConsumerMessage) {
	partitionConsumer, err := consumer.ConsumePartition(topics[0], partition, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	channel := partitionConsumer.Messages()
	for msg := range channel {
		msgChan <- msg
	}
}

func Consumer(topics []string, servers string, msgChan chan *sarama.ConsumerMessage) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{servers}, config)
	if err != nil {
		panic(err)
	}
	partitions, err := consumer.Partitions(topics[0])
	if err != nil {
		panic(err)
	}
	if len(partitions) == 0 {
		panic("No partitions found")
	}
	for _, partition := range partitions {
		consumePartition(consumer, topics, partition, msgChan)
	}
}
