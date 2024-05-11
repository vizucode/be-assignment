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

	return db
}
