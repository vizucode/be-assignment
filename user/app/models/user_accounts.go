package models

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	UserId      uint
	AccountId   uint
	AccountName string
	AccountCode string `gorm:"unique"`

	User Users `gorm:"foreignKey:UserId;references:ID"`

	Transactions []Transaction `gorm:"foreignKey:AccountCodeOwner;references:AccountCode"`
}
