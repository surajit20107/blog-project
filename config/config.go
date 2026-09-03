package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DATABASE_URL string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DATABASE_URL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/blog_project?sslmode=disable"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
