package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/livingdolls/go-template/internal/adapter/http"
	"github.com/livingdolls/go-template/internal/bootstrap"
	"github.com/livingdolls/go-template/internal/core/email"
	"github.com/livingdolls/go-template/internal/core/events"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/messagebroker"
	"go.uber.org/zap"
)

type NotificationHandler struct{}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

func (h *NotificationHandler) Handle(ctx context.Context, msg []byte) error {
	// Logika pengolahan pesan notification
	fmt.Printf("Processing notification message: %s\n", msg)
	return nil
}

func main() {
	// Inisialisasi aplikasi
	db := bootstrap.Setup()
	defer db.Close()
	defer logger.SyncLogger()

	registry := messagebroker.NewHandlerRegistry()

	registry.RegisterHandler(events.EmailVerificationEvent, email.NewEmailHandler())
	registry.RegisterHandler(events.NotificationSendEvent, NewNotificationHandler())

	rabbitMQ, err := messagebroker.NewRabbitMQAdapter("amqp://guest:guest@rabbitmq:5672/", registry)
	if err != nil {
		logger.Log.Fatal("Failed to connect to RabbitMQ: ", zap.Error(err))
	}
	defer rabbitMQ.Close()

	// Jalankan Consumer untuk mendengarkan pesan dari queue
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Print("jalan")
		err := rabbitMQ.Consume(ctx, "event_queue", "events", events.EmailVerificationEvent)
		if err != nil {
			log.Fatalf("Failed to start consumer: %v", err)
		}
	}()

	go func() {
		err := rabbitMQ.Consume(ctx, "event_queue", "events", events.NotificationSendEvent) // Untuk notification
		if err != nil {
			log.Fatalf("Failed to start consumer: %v", err)
		}
	}()

	// Simulasi Publish Event
	time.Sleep(2 * time.Second) // Tunggu RabbitMQ siap
	err = rabbitMQ.Publish(ctx, "events", events.EmailVerificationEvent, map[string]string{"ss": "1212"})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Simulasi Publish Event untuk notification
	err = rabbitMQ.Publish(ctx, "events", events.NotificationSendEvent, map[string]string{"message": "This is a notificatison"})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Tunggu agar pesan bisa diproses sebelum aplikasi berhenti
	time.Sleep(5 * time.Second)

	// Mulai server
	server := server.StartServer(db)

	// Menunggu sinyal shutdown
	waitForShutdown(server)
}

// waitForShutdown menangani shutdown aplikasi dengan gracefull
func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	// Menunggu sinyal shutdown
	sig := <-quit
	logger.Log.Info("Received shutdown signal", zap.String("signal", sig.String()))

	// Menutup server dengan gracefull shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error("Server forced to shutdown", zap.Error(err))
	} else {
		logger.Log.Info("Server shutdown gracefully")
	}
}
