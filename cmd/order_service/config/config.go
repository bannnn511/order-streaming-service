package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"order-streaming-services/pkg/config"
	"os"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
		Kafka        `yaml:"kafka"`
	}

	Kafka struct {
		URL string `env-required:"true" yaml:"url" env:"KAFKA_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println("config path: " + dir)

	err = cleanenv.ReadConfig(dir+"/cmd/order_service/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
