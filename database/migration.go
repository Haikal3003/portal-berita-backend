package database

import (
	"log"

	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	errDrop := DB.Migrator().DropTable(
		&models.Comment{},
		&models.Notification{},
		&models.Tag{},
		&models.Category{},
		&models.Article{},
		&models.Profile{},
		&models.User{},
	)
	if errDrop != nil {
		log.Println("Warning: Failed to drop tables:", errDrop)
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
		&models.Comment{},
	)

	if errMigrate != nil {
		log.Fatal("Failed to migrate database: ", errMigrate)
	}

	log.Println("✅ Database migrated successfully")
}
