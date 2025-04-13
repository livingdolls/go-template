package port

type EventPublisher interface {
	Publish(eventName string, payload []byte) error
}
