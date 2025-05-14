package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

<<<<<<< HEAD
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: false})
=======
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
>>>>>>> f7757e78e15bac05ea35d7d2c9ac267aa5db3a4b
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	DB = db

	log.Println("Database connected successfully")

}
