package logger

import (
	"os"

	"github.com/livingdolls/go-template/internal/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(config config.AppConfig) {
	// Konversi level log dari string ke zapcore.Level
	logLevel := zapcore.InfoLevel
	if lvl, err := zapcore.ParseLevel(config.Log.Level); err == nil {
		logLevel = lvl
	}

	// Konfigurasi encoder log
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Logger untuk file INFO
	infoWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Log.Files.Info,
		MaxSize:    config.Log.MaxSize,
		MaxBackups: config.Log.MaxBackups,
		MaxAge:     config.Log.MaxAge,
		Compress:   config.Log.Compress,
	})

	// Logger untuk file ERROR
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Log.Files.Error,
		MaxSize:    config.Log.MaxSize,
		MaxBackups: config.Log.MaxBackups,
		MaxAge:     config.Log.MaxAge,
		Compress:   config.Log.Compress,
	})

	// Logger untuk file APP (semua log)
	appWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "app.log", // Nama file untuk semua log
		MaxSize:    config.Log.MaxSize,
		MaxBackups: config.Log.MaxBackups,
		MaxAge:     config.Log.MaxAge,
		Compress:   config.Log.Compress,
	})

	// Encoder JSON
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Console Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Core untuk masing-masing log
	infoCore := zapcore.NewCore(jsonEncoder, infoWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	}))

	errorCore := zapcore.NewCore(jsonEncoder, errorWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	}))

	appCore := zapcore.NewCore(jsonEncoder, appWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true // Menyimpan semua level log
	}))

	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	// Gabungkan semua core
	core := zapcore.NewTee(infoCore, errorCore, appCore, consoleCore)

	// Inisialisasi logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// Set sebagai logger global
	zap.ReplaceGlobals(Log)
}

// Pastikan dipanggil saat aplikasi shutdown
func SyncLogger() {
	if Log != nil {
		_ = Log.Sync()
	}
}
