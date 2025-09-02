package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load loads environment variables from .env file
func Load() error {
	// Load .env file in non-production environments
	env := os.Getenv("APP_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			fmt.Printf("Warning: .env file not found: %v\n", err)
			// Don't return error - environment variables might be passed via Docker
		}
	}
	return nil
}

// GetEnv gets an environment variable with a default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
