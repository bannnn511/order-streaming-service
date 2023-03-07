package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"order-streaming-services/cmd/order_service/config"
	"order-streaming-services/internal/order_service/app"
	"order-streaming-services/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	slog.Info("⚡ init app", "name", cfg.Name, "version", cfg.Version)

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	app.Run(ctx, cancel, cfg)

	select {
	case <-quit:
		slog.Info("signal.Notify")
	case <-ctx.Done():
		slog.Info("ctx.Done()", "app Done")
	}
}
