package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	UserId            int     `gorm:"not null"`
	AccountCodeOwner  string  `gorm:"not null"`
	TargetAccountCode string  `gorm:"not null"`
	TransactionType   string  `gorm:"not null"`
	IncomSpending     string  `gorm:"not null"`
	CurrencyType      string  `gorm:"not null"`
	Amount            float64 `gorm:"not null"`
	Status            bool    `gorm:"not null"`

	AccountOwner  UserAccount `gorm:"foreignKey:AccountCodeOwner;references:AccountCode"`
	AccountTarget UserAccount `gorm:"foreignKey:TargetAccountCode;references:AccountCode"`
}
