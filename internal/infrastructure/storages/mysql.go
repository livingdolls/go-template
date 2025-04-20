package storages

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/livingdolls/go-template/internal/config"
	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	Database *sqlx.DB
}

// Close implements port.DatabasePort.
func (d *database) Close() error {
	if err := d.Database.Close(); err != nil {
		logger.Log.Fatal("error closing database", zap.Error(err))
		return fmt.Errorf("error closing database: %w", err)
	}
	return nil
}

// GetDatabase implements port.DatabasePort.
func (d *database) GetDatabase() *sqlx.DB {
	return d.Database
}

func NewDatabase(config config.DatabaseConfig) (port.DatabasePort, error) {
	db, err := openDatabase(config)

	if err != nil {
		logger.Log.Fatal("failed to connect database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect database %v", err)
	}

	logger.Log.Info("database connection established successfully")
	return &database{
		Database: db,
	}, nil
}

func openDatabase(config config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sqlx.Open(config.Driver, dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenCons)
	db.SetMaxIdleConns(config.MaxIdleCons)
	db.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Minute)

	if err := db.Ping(); err != nil {
		logger.Log.Error("database ping error", zap.Error(err))
		return nil, fmt.Errorf("database ping error: %w", err)
	}

	logger.Log.Info("database connected successfully")
	return db, nil
}
