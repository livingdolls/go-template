package main

import (
	"log"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
)

func main() {
	config, err := config.LoadConfig("config")

	if err != nil {
		log.Fatalf("failed to load configuration file %v", err)
	}

	logger.InitLogger(config)
	defer logger.SyncLogger()

	logger.Log.Info("aplikasi di mulai")
	logger.Log.Error("error disini")
}
