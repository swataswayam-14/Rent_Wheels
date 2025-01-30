package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'user'"`
}
type Vehicle struct {
	gorm.Model
	Make         string  `gorm:"not null"`
	ModelName    string  `gorm:"not null"`
	Year         int     `gorm:"not null"`
	LicensePlate string  `gorm:"unique;not null"`
	DailyRate    float64 `gorm:"not null"`
	Status       string  `gorm:"default:'available'"`
	Bookings     []Booking
}

type Booking struct {
	gorm.Model
	UserID    int       `gorm:"not null"`
	VehicleID int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	TotalCost float64   `gorm:"not null"`
	Status    string    `gorm:"default:'pending'"`
	PaymentID string
}
