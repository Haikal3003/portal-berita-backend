package main

import (
	"log"
	"os"
	"portal-berita-backend/database"
	"portal-berita-backend/handlers"
	"portal-berita-backend/routes"
	"portal-berita-backend/seed"
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("DB_DSN:", os.Getenv("DB_DSN"))

	database.ConnectDatabase()
	database.AutoMigrateTables()
	seed.SetupAdmin()

	app := fiber.New()

	// init service dan handler
	authService := services.NewAuthService(database.DB)
	authHandler := handlers.NewAuthHandler(authService)

	userService := services.NewUserService(database.DB)
	userHandler := handlers.NewUserHandler(userService)

	// init routes
	api := app.Group("/api")
	routes.AuthRoutes(api, authHandler)
	routes.UserRoutes(api, userHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
