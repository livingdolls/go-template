package config

import (
	"log"

	"github.com/spf13/viper"
)

type LoggerFiles struct {
	Info  string
	Error string
}

type LoggerConfig struct {
	Level      string
	Files      LoggerFiles
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type AppConfig struct {
	Log LoggerConfig
}

var Config AppConfig

func LoadConfig(path string) (config AppConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error read configurasi file %v", err)
	}

	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("error mengurai configurasi %v", err)
	}

	return
}
