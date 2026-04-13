package config

import (
	"log"
	"os"
)

type Config struct {
	AppEnv      string
	HTTPPort    string
	DatabaseURL string
	Issuer      string
	Secret      string
}

func MustLoad() Config {
	cfg := Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		HTTPPort:    getEnv("HTTP_PORT", "8085"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres123@localhost:5432/realestate?sslmode=disable"),
		Issuer:      getEnv("JWT_ISSUER", "TordiaTech"),
		Secret:      getEnv("JWT_SECRET", "7u2y67819172hwy265u8y"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
