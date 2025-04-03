package messagebroker

import (
	"context"
	"encoding/json"

	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RabbitMQAdapter struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	registry  port.MessageTypeRegistry
	exchanges map[string]bool
	queues    map[string]bool
}

func NewRabbitMQAdapter(url string, registry port.MessageTypeRegistry) (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Log.Error("failed to connect rabbitmq", zap.Error(err))
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Log.Error("failed to init channel", zap.Error(err))
		return nil, err
	}

	return &RabbitMQAdapter{
		conn:      conn,
		channel:   ch,
		registry:  registry,
		exchanges: make(map[string]bool),
		queues:    make(map[string]bool),
	}, nil
}

func (r *RabbitMQAdapter) ensureExchange(exchange string) error {
	if !r.exchanges[exchange] {
		err := r.channel.ExchangeDeclare(
			exchange,
			"topic", // menggunakan topic untuk routing fleksibel,
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,
		)
		if err != nil {
			logger.Log.Error("failed to declare exchange", zap.Error(err))
			return err
		}
		r.exchanges[exchange] = true
	}
	return nil
}

func (r *RabbitMQAdapter) Publish(ctx context.Context, exchange, routingKey string, payload interface{}) error {
	if err := r.ensureExchange(exchange); err != nil {
		logger.Log.Error("ensure exchange error", zap.Error(err))
		return err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Error("failed to marshal payload json", zap.Error(err))
		return err
	}

	return r.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
			Headers: amqp.Table{
				"event_type": routingKey,
			},
		},
	)
}

func (r *RabbitMQAdapter) Consume(ctx context.Context, queue, exchange, routingKey string) error {
	// Pastikan exchange sudah dideklarasikan sebelum digunakan
	if err := r.ensureExchange(exchange); err != nil {
		logger.Log.Error("Failed to ensure exchange", zap.String("exchange", exchange), zap.Error(err))
		return err
	}

	// Deklarasi queue
	_, err := r.channel.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		logger.Log.Error("Failed to declare queue", zap.String("queue", queue), zap.Error(err))
		return err
	}

	// Bind queue ke exchange dengan routing key
	err = r.channel.QueueBind(
		queue,
		routingKey,
		exchange,
		false, // noWait
		nil,   // args
	)
	if err != nil {
		logger.Log.Error("Failed to bind queue", zap.String("queue", queue), zap.String("exchange", exchange), zap.Error(err))
		return err
	}

	msgs, err := r.channel.Consume(
		queue,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)

	if err != nil {
		logger.Log.Error("Failed to consume messages", zap.String("queue", queue), zap.Error(err))
		return err
	}

	logger.Log.Info("queue receuver", zap.String("queue name", queue))

	// Proses pesan
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgs:
				eventType, ok := msg.Headers["event_type"].(string)
				if !ok {
					msg.Nack(false, false)
					continue
				}

				logger.Log.Info("queue receuver", zap.String("events", eventType))

				handler, exists := r.registry.GetHandler(eventType)

				if !exists {
					msg.Nack(false, false)
					continue
				}

				if err := handler.Handle(ctx, msg.Body); err != nil {
					msg.Nack(false, true) // requeue
				} else {
					msg.Ack(false)
				}
			}
		}
	}()

	return nil
}

func (r *RabbitMQAdapter) Close() error {
	var err error

	if r.channel != nil {
		err = r.channel.Close()
		if err != nil {
			logger.Log.Error("failed to close channel", zap.Error(err))
		}
	}

	if r.conn != nil {
		err = r.conn.Close()
		if err != nil {
			logger.Log.Error("failed to close connection", zap.Error(err))
		}
	}

	return err
}
