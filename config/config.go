package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Env         string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	_ = godotenv.Load()

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		Env:         os.Getenv("ENV"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.Env == "" {
		config.Env = "development"
	}

	if config.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return config, nil
}
