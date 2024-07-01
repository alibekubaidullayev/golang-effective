package db

import (
	"log/slog"
	"rest/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(AppConfig core.Config) {
	slog.Info("Initializing DataBase")
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to db")
	}
	err = DB.AutoMigrate(&Person{})
	if err != nil {
		slog.Error("Failed to apply migrations: %v", err)
	}
}
