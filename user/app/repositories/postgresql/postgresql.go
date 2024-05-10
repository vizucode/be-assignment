package postgresql

import "gorm.io/gorm"

type postgre struct {
	db *gorm.DB
}

func NewDatabase(db *gorm.DB) *postgre {
	return &postgre{db: db}
}
