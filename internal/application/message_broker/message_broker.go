package message_broker

type MessageBroker interface {
	Consumer(params interface{}, msgChan chan string)
	Producer(params interface{}, msgChan chan string)
}
