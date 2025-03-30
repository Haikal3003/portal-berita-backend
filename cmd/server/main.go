package main

import (
	"os"
	"portal-berita-backend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	database.AutoMigrateTables()

	app := fiber.New()

	app.Listen(":" + os.Getenv("PORT"))
}
