package messagebroker

import (
	"context"

	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type RabbitMQPublisher struct {
	rmq *RabbitMQAdapter
	ctx context.Context
}

func NewRabbitMQPublisher(rmq *RabbitMQAdapter) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		rmq: rmq,
		ctx: context.Background(),
	}
}

func (p *RabbitMQPublisher) Publish(eventName string, payload []byte) error {
	logger.Log.Info("publish event", zap.String("event Name", eventName))
	return p.rmq.Publish(p.ctx, "events", eventName, payload)
}
