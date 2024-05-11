package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountType   string  `gorm:"not null"`
	LimitBallance float64 `gorm:"not null"`
	Status        bool

	UserAccount []UserAccount `gorm:"foreignKey:AccountId"`
}
