package db

import (
	"log/slog"
	"rest/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(AppConfig core.Config) error {
	slog.Info("Initializing DataBase")

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return err
	}

	if err := applyMigrations(); err != nil {
		return err
	}

	slog.Info("Database initialized successfully")
	return nil
}

func applyMigrations() error {
	slog.Info("Applying migrations")
	if err := DB.AutoMigrate(&Person{}); err != nil {
		slog.Error("Failed to apply Person migrations", "error", err)
		return err
	}
	return nil
}
