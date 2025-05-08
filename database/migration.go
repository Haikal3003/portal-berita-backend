package database

import (
	"log"
	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	tables := []interface{}{
		&models.User{},
		&models.Profile{},
		&models.Article{},
		&models.Category{},
		&models.ArticleCategory{},
		&models.Tag{},
		&models.ArticleTag{},
		&models.Comment{},
		&models.Notification{},
	}

	for _, table := range tables {
		if !DB.Migrator().HasTable(table) {
			if err := DB.AutoMigrate(table); err != nil {
				log.Fatalf("Failed to migrate table: %T | Error: %v", table, err)
			} else {
				log.Printf("Migrated table: %T\n", table)
			}
		} else {
			log.Printf("ℹTable already exists: %T\n", table)
		}
	}

	log.Println("Database migration completed successfully")

}
