package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/storages"
	"go.uber.org/zap"
)

func main() {
	// Inisialisasi aplikasi
	db := initialize()
	defer db.Close()
	defer logger.SyncLogger()

	// Mulai server
	server := StartServer(db)

	// Menunggu sinyal shutdown
	waitForShutdown(server)
}

// initialize melakukan setup awal aplikasi
func initialize() port.DatabasePort {
	// Load configuration
	if err := config.LoadConfig("config"); err != nil {
		zap.L().Fatal("Failed to load configuration file", zap.Error(err))
	}

	// Initialize logger
	logger.InitLogger(config.Config)

	// Initialize database
	db, err := storages.NewDatabase(config.Config.Database)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	logger.Log.Info("Application successfully initialized")
	return db
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
