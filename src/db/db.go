package db

import (
	"log/slog"
	"reflect"
	"rest/core"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(AppConfig core.Config) error {
	slog.Info("Initializing DataBase")

	var err error

	dsn := ("host=" + AppConfig.Host +
		" user=" + AppConfig.User +
		" password=" + AppConfig.Password +
		" dbname=" + AppConfig.DBName +
		" port=" + AppConfig.DBPort +
		" sslmode=disable")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

	models := []interface{}{
		&Person{},
		&Task{},
		&TaskUser{},
	}

	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			slog.Error(
				"Failed to apply migrations for",
				"model",
				reflect.TypeOf(model).Elem().Name(),
			)
			return err
		}
	}

	return nil
}
