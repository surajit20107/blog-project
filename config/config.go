package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DbHost string
	DbUser string
	DbPassword string
	DbName string
	DbPort string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "postgres"),
		DbName:     getEnv("DB_NAME", "blogdb"),
		DbPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
