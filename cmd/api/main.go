package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/livingdolls/go-template/internal/adapter/http"
	"github.com/livingdolls/go-template/internal/bootstrap"
	_ "github.com/livingdolls/go-template/internal/core/email"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/messagebroker"
	"go.uber.org/zap"
)

func main() {
	// Inisialisasi aplikasi
	db := bootstrap.Setup()
	defer db.Close()
	defer logger.SyncLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rmq, err := messagebroker.NewRabbitMQAdapter("amqp://guest:guest@rabbitmq:5672", messagebroker.GlobalRegistry())
	if err != nil {
		logger.Log.Fatal("failed to start rabbitmq", zap.Error(err))
	}

	// Start RabbitMQ consumer
	if err := messagebroker.StartRabbitMQConsumer(ctx, rmq); err != nil {
		logger.Log.Fatal("Failed to start rabbitmq consumer", zap.Error(err))
	}

	server := server.StartServer(db, rmq)

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
