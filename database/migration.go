package database

import (
	"log"
	"portal-berita-backend/models"
)

func AutoMigrateTables() {
	if DB == nil {
		log.Fatal("Database is not connected")
	}

	tables := []interface{}{
		&models.Bookmark{},
		&models.Like{},
		&models.Comment{},
		&models.Notification{},
		&models.Tag{},
		&models.Category{},
		&models.Article{},
		&models.Profile{},
		&models.User{},
	}

	for _, table := range tables {
		errDrop := DB.Migrator().DropTable(table)
		if errDrop != nil {
			log.Println("Warning: Failed to drop table:", errDrop)
		}
	}

	errMigrate := DB.AutoMigrate(
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
	if errMigrate != nil {
		log.Fatal("Failed to migrate database: ", errMigrate)
	}

	log.Println("Database migrated successfully")
}
