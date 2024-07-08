package main

import (
	"log/slog"
	"os"

	"rest/core"
	"rest/db"
	"rest/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginlogger "github.com/FabienMht/ginslog/logger"
	ginrecovery "github.com/FabienMht/ginslog/recovery"
	"github.com/phsym/console-slog"
)

var AppConfig *core.Config

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file.")
		os.Exit(1)
	}
	slog.Info("Starting Application")

	AppConfig = core.LoadConfig()

	if err := db.InitDB(*AppConfig); err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}

	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{Level: slog.LevelDebug}),
	)
	slog.SetDefault(logger)

	router := gin.New()
	router.Use(ginlogger.New(logger))
	router.Use(ginrecovery.New(logger))
	routes.RegisterUserRoutes(router, "users")
	routes.RegisterTaskRoutes(router, "tasks")

	port := AppConfig.Port
	slog.Info("Starting server on", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
