package bootstrap

import (
	"path/filepath"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"github.com/livingdolls/go-template/internal/infrastructure/storages"
	"go.uber.org/zap"
)

func Setup() port.DatabasePort {
	configPath, err := filepath.Abs("config")

	if err != nil {
		zap.L().Fatal("Gagal mendapatkan path absolute", zap.Error(err))
	}

	// Load configuration
	if err := config.LoadConfig(configPath); err != nil {
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
