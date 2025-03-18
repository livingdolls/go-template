package config

import (
	"github.com/spf13/viper"
)

type LoggerFiles struct {
	Info  string
	Error string
}

type DatabaseConfig struct {
	Driver      string
	Host        string
	Port        string
	User        string
	Password    string
	Name        string
	MaxOpenCons int
	MaxIdleCons int
	MaxLifeTime int
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
	Log      LoggerConfig
	Database DatabaseConfig
}

var Config AppConfig

func LoadConfig(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := viper.Unmarshal(&Config)

	if err != nil {
		return err
	}

	return nil
}
