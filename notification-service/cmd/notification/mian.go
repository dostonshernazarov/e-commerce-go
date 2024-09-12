package main

import (
	"ekzamen_5/notification-service/internal/app"
	"ekzamen_5/notification-service/internal/config"
	"ekzamen_5/notification-service/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	log.Info("Starting service1", slog.Any(
		"config", cfg.RPCPort))
	application := app.NewApp(cfg, log)

	go application.GrpcApp.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("received shutdown signal", slog.String("signal", sig.String()))
	application.GrpcApp.Stop()
	log.Info("shutting down server")
}
