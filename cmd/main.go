package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/storages"
	"go.uber.org/zap"
)

func main() {
	db := initialize()
	defer db.Close()

	waitForShutdown()
}

func initialize() port.DatabasePort {
	// Load configuration
	if err := config.LoadConfig("config"); err != nil {
		zap.L().Fatal("Failed to load configuration file", zap.Error(err))
	}

	// Initialize logger
	logger.InitLogger(config.Config)
	defer logger.SyncLogger()

	// Initialize database
	db, err := storages.NewDatabase(config.Config.Database)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	logger.Log.Info("Application successfully initialized")
	return db
}

func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	sig := <-quit
	logger.Log.Info("Received shutdown signal", zap.String("signal", sig.String()))
	logger.Log.Info("Graceful shutdown initiated")
}
