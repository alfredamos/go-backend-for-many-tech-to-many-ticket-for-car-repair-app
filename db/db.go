package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	//----> Connect gorm to mysql database
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//----> Check for error.
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	return DB, nil
}
