package models

import (
	"gorm.io/gorm"
)

type Query struct {
	gorm.Model
	Path       string `gorm:"not null;"`
	RemoteAddr string `gorm:"not null;"`
	FizzBuzzId uint
}
