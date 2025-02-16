package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// getEnv gets the value of an environment variable or returns a fallback
func getEnv(key, fallback string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return env
}

// getJwtDuration gets the duration of the JWT token
func getJwtDuration(fallback string) time.Duration {
	duration := getEnv("JWT_DURATION", fallback)
	jwtDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalf("Invalid duration: %s", duration)
	}
	return jwtDuration
}
