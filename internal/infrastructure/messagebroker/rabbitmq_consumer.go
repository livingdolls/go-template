package messagebroker

import (
	"context"

	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/pkg/events"
	"go.uber.org/zap"
)

func StartRabbitMQConsumer(ctx context.Context, rmq *RabbitMQAdapter) error {
	go func() {
		err := rmq.Consume(ctx, "event_queue", "events",
			events.EmailVerificationEvent,
		)

		logger.Log.Info("starting consume events")

		if err != nil {
			logger.Log.Fatal("Rabbitmq consume error:", zap.Error(err))
		}
	}()

	return nil
}
