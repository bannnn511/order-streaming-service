package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"order-streaming-services/cmd/orders/config"
	"order-streaming-services/internal/orders/app"
	"order-streaming-services/pkg/logger"
	"os"
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

	app.Run(cfg)
}
