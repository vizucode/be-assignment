package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email     string `gorm:"unique"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Status    bool
}
