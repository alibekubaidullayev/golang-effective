package core

import (
	"os"
)

func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "1"),
		DBName:   getEnv("DB_NAME", "rest"),
		DBPort:   getEnv("DB_PORT", "5432"),
	}
}

type Config struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
	DBPort   string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
