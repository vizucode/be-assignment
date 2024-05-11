package models

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	UserId      uint
	AccountId   uint
	AccountName string
	AccountCode string
}
