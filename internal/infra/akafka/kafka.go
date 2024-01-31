package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consumer(topics []string, servers string, msgChan chan *kafka.Message) {
}
