package main

import (
	"log/slog"
	"net/http"

	"rest/core"
	"rest/db"
	"rest/routes"

	"github.com/joho/godotenv"
)

var AppConfig *core.Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file. Default config values will be used")
	}
	AppConfig = core.LoadConfig()
}

func main() {
	slog.Info("Starting Application")

	db.InitDB(*AppConfig)
	router := http.NewServeMux()
	routes.RegisterUserRoutes(router, "usersas")
	port := AppConfig.Port

	slog.Info("Starting server on", "port", port)
	slog.Debug("Debuging")

	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("Could not start server: %v", err)
	}
}
