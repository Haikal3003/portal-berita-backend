package database

import (
	"log"
	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database is not connected")
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Article{},
		&models.Category{},
		&models.Tag{},
		&models.Notification{},
		&models.Comment{},
		&models.Like{},
		&models.Bookmark{},
	)

	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")
}
