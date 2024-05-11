package database

import (
	"log"
	"user_service/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {

	// dbhost := os.Getenv("DB_HOST")
	// dbname := os.Getenv("DB_NAME")
	// dbusername := os.Getenv("DB_USERNAME")
	// dbpassword := os.Getenv("DB_PASSWORD")
	// dbport := os.Getenv("DB_PORT")

	dbhost := "aws-0-ap-southeast-1.pooler.supabase.com"
	dbname := "postgres"
	dbusername := "postgres.qytwpijdbksagnkuqgay"
	dbpassword := "DMLzDUfRz8M4L9h4"
	dbport := "5432"

	dsn := "host=" + dbhost + " user=" + dbusername + " password=" + dbpassword + " dbname=" + dbname + " port=" + dbport + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Users{},
	)

	return db
}
