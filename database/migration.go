package database

import (
	"log"
	"strings"

	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("Failed to create uuid extension: %v", err)
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Article{},
		&models.SavedArticle{},
		&models.Category{},
		&models.Tag{},
		&models.Comment{},
		&models.Notification{},
	)

	if err != nil {
		// Tangani error spesifik
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Table already exists, skipping migration.")
		} else if strings.Contains(err.Error(), "prepared statement") {
			log.Println("Prepared statement already exists. Consider disabling statement cache or resetting connection.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Database migration completed successfully.")
	}
}
