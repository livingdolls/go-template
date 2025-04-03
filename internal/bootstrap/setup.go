package bootstrap

import (
	"os"
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
		println("Gagal mendapatkan path absolute:", err.Error()) // Gunakan `println` sementara
		os.Exit(1)
	}

	// Load configuration
	if err := config.LoadConfig(configPath); err != nil {
		println("Failed to load configuration file:", err.Error()) // Gunakan `println` sementara
		os.Exit(1)
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

// func setupMessageBroker() (port.MessagePublisher, port.MessageConsumer, error) {
// 	if err := loadConfig(); err != nil {
// 		return nil, nil, err
// 	}

// 	// Initialize RabbitMQ
// 	rabbitMQUrl := fmt.Sprintf("amqp://%v:%v@%v:%v/",
// 		config.Config.RabbitMQ.User,
// 		config.Config.RabbitMQ.Password,
// 		config.Config.RabbitMQ.Host,
// 		config.Config.RabbitMQ.Port,
// 	)

// 	registry := messagebroker.NewHandlerRegistry()

// 	adapter, err := messagebroker.NewRabbitMQAdapter(rabbitMQUrl, registry)
// }

// func loadConfig() error {
// 	configPath, err := filepath.Abs("config")

// 	if err != nil {
// 		println("Gagal mendapatkan path absolute:", err.Error()) // Gunakan `println` sementara
// 		os.Exit(1)
// 	}

// 	if err := config.LoadConfig(configPath); err != nil {
// 		println("Failed to load configuration file:", err.Error()) // Gunakan `println` sementara
// 		os.Exit(1)
// 	}

// 	logger.InitLogger(config.Config)
// 	return nil
// }
