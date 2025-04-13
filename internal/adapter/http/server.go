package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/messagebroker"
	"go.uber.org/zap"
)

func StartServer(db port.DatabasePort, rmq *messagebroker.RabbitMQAdapter) *http.Server {
	rmqPublisher := messagebroker.NewRabbitMQPublisher(rmq)

	router := SetupRouter(db, rmqPublisher)

	serverHost := fmt.Sprintf(":%v", config.Config.Server.Port)

	server := &http.Server{
		Addr:    serverHost,
		Handler: router,
	}

	logger.Log.Info("Starting server", zap.String("host", serverHost))

	// Jalankan server dalam goroutine agar tidak blocking
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server error", zap.Error(err))
		}
	}()

	return server
}

// ShutdownServer menangani proses shutdown server dengan baik
func ShutdownServer(server *http.Server) {
	// Gunakan timeout agar tidak langsung mematikan server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error("Server forced to shutdown", zap.Error(err))
	} else {
		logger.Log.Info("Server shutdown gracefully")
	}
}
