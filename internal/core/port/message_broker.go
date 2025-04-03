package port

import "context"

// MessagePublisher menangani berbagai jenis pesan
type MessagePublisher interface {
	Publish(ctx context.Context, exchange string, routingKey string, payload interface{}) error
	Close() error
}

// MessageConsumer menangani konsumsi berbagai pesan
type MessageConsumer interface {
	Consume(ctx context.Context, queue string, handler func(payload []byte) error) error
	Close() error
}

// MessageTypeRegistry untuk routing message handler
type MessageTypeRegistry interface {
	RegisterHandler(messageType string, handler MessageHandler)
	GetHandler(messageType string) (MessageHandler, bool)
	UnregisterHandler(messageType string)
	GetAllHandlers() map[string]MessageHandler
}

type MessageHandler interface {
	Handle(ctx context.Context, payload []byte) error
}
