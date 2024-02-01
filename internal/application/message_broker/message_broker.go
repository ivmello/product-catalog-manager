package message_broker

type MessageBroker interface {
	Consumer(destinations []string, msgChan chan string) error
	Producer(msg string, destinations []string) error
}
