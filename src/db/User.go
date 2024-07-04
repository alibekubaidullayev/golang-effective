package db

import "time"

type Person struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PassNumber string
	Surname    string
	Name       string
	Patronymic string
	Address    string
}

func (Person) TableName() string {
	return "person"
}
