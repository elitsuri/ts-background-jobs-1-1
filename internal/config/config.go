package config

import (
	"os"
	"strconv"
)

// Config holds all application configuration.
type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	Environment string
	JWTExpiry   int // minutes
}

// Load reads configuration from environment variables with sane defaults.
func Load() *Config {
	expiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_MINUTES", "60"))
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/ts-background-jobs-1?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "change-this-to-a-secret-key-min-32-chars"),
		Environment: getEnv("ENVIRONMENT", "development"),
		JWTExpiry:   expiry,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
