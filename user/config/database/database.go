package database

import (
	"log"
	"os"
	"user_service/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {

	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	dbusername := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")

	dsn := "host=" + dbhost + " user=" + dbusername + " password=" + dbpassword + " dbname=" + dbname + " port=" + dbport + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Users{},
		&models.Account{},
		&models.UserAccount{},
		&models.Transaction{},
	)

	var isDebit int64
	var isCredit int64
	var isLoan int64
	db.Model(&models.Account{}).Where("account_type = ?", "DEBIT").Count(&isDebit)
	db.Model(&models.Account{}).Where("account_type = ?", "CREDIT").Count(&isCredit)
	db.Model(&models.Account{}).Where("account_type = ?", "LOAN").Count(&isLoan)

	if isDebit < 1 {
		db.Create(&models.Account{AccountType: "DEBIT", LimitBallance: 0, Status: true})
	}
	if isCredit < 1 {
		db.Create(&models.Account{AccountType: "CREDIT", LimitBallance: 0, Status: true})
	}
	if isLoan < 1 {
		db.Create(&models.Account{AccountType: "LOAN", LimitBallance: 0, Status: true})
	}

	return db
}
