package db

import (
	"errors"
	"fmt"
	"time"
)

type Status string

type FieldLength struct {
	Field  string
	Length int
}

type Validatable interface {
	Validate() error
}

const (
	Pending    Status = "pending"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
)

type Person struct {
	ID         uint   `gorm:"primaryKey"`
	PassNumber string `gorm:"not null;unique;size:10"`
	Surname    string `gorm:"not null;size:100"`
	Name       string `gorm:"not null;size:100"`
	Patronymic string `gorm:"not null;size:100"`
	Address    string `gorm:"not null;size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Task struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Status      Status     `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   time.Time  `gorm:"not null"`
	DueDate     time.Time  `gorm:"not null"`
	CompletedAt *time.Time `gorm:"default:null"`
}

type TaskUser struct {
	ID          uint   `gorm:"primaryKey"`
	TaskID      uint   `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	Task        Task   `gorm:"foreignKey:TaskID"`
	User        Person `gorm:"foreignKey:UserID"`
	StartDate   time.Time
	EndDate     time.Time
	PaymentRate float64 `gorm:"not null"`
}

func (p Person) Validate() error {
	fields := map[string]FieldLength{
		"Surname":    {p.Surname, 100},
		"Name":       {p.Name, 100},
		"Patronymic": {p.Patronymic, 100},
		"Address":    {p.Address, 100},
	}

	for fieldName, fieldValue := range fields {
		if len(fieldValue.Field) > fieldValue.Length {
			return errors.New(
				fmt.Sprintf(
					"%s exceeds %d characters",
					fieldName,
					fieldValue.Length),
			)
		}
	}

	return nil
}
