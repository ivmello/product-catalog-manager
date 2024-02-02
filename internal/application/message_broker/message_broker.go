package message_broker

type MessageBroker interface {
	Consumer(params interface{}, msgChan chan []byte)
	Producer(params interface{}, msgChan chan []byte)
}
