package database

import (
	"log"

	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Article{},
		&models.SavedArticle{},
		&models.Category{},
		&models.Tag{},
		&models.Comment{},
		&models.Notification{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	} else {
		log.Println("✅ All tables migrated successfully")
	}

}
