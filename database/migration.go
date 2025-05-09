package database

import (
	"log"
	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	errMigrate := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Article{},
		&models.SavedArticle{},
		&models.Category{},
		&models.Tag{},
		&models.Comment{},
		&models.Notification{},
	)

	if errMigrate != nil {
		log.Fatal("migration fail")
	}

	log.Println("Database migration completed successfully")

}
