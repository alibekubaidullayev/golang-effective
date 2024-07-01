package db

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	PassSeries uint
	PassNumber uint
	Surname    string
	Name       string
	Patronymic string
	Address    string
}

func (Person) TableName() string {
	return "person"
}
