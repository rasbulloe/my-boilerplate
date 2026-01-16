package models

import "time"

type User struct {
	ID         int64 `gorm:"primaryKey;"`
	Name       string
	Email      string
	Password   string
	Address    string
	Phone      string
	Photo      string
	Latitude   string
	Longitude  string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
