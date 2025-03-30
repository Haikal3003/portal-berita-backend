package database

import "log"

func AutoMigrateTables() {
	if DB != nil {
		log.Fatal("Database is not connected")
	}

	err := DB.AutoMigrate()

	if err != nil {
		log.Fatal("failed to migrate database")
	}

	log.Println("Database migrated successfully")

}
