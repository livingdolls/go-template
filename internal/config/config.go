package config

import (
	"github.com/spf13/viper"
)

type LoggerFiles struct {
	Info   string
	Error  string
	Global string
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

type RabbitMQConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type LoggerConfig struct {
	Level      string
	Files      LoggerFiles
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type ServerConfig struct {
	Port int
}

type AppConfig struct {
	Log      LoggerConfig
	Database DatabaseConfig
	RabbitMQ RabbitMQConfig
	SMTP     SMTPConfig
	Server   ServerConfig
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
