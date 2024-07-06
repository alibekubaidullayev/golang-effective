package core

import (
	"fmt"
	"os"
)

func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Host:     getEnv("HOST", "localhost"),
		User:     getEnv("USER", "postgres"),
		Password: getEnv("PASSWORD", "1"),
		DBName:   getEnv("DBNAME", "rest"),
		DBPort:   getEnv("DBPORT", "5432"),
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

func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.DBPort,
	)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
