package db

import (
	"log/slog"
	"reflect"
	"rest/core"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(AppConfig core.Config) error {
	slog.Info("Initializing DataBase")

	var err error
	// dsn := "host=localhost user=user password=1 dbname=db port=5432"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
