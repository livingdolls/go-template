package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/livingdolls/go-template/internal/adapter/http"
	"github.com/livingdolls/go-template/internal/bootstrap"
	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func main() {
	// Inisialisasi aplikasi
	db := bootstrap.Setup()
	defer db.Close()
	defer logger.SyncLogger()

	// Mulai server
	server := server.StartServer(db)

	// Menunggu sinyal shutdown
	waitForShutdown(server)
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
