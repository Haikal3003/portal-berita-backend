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

	articleService := services.NewArticleService(database.DB)
	articleHandler := handlers.NewArticleHandler(articleService)

	// init routes
	api := app.Group("/api")
	routes.AuthRoutes(api, authHandler)
<<<<<<< HEAD
	routes.ProfileRoutes(api, profileHandler)
	routes.ArticleRoutes(api, articleHandler)
=======
	routes.UserRoutes(api, userHandler)
>>>>>>> f7757e78e15bac05ea35d7d2c9ac267aa5db3a4b

	app.Listen(":" + os.Getenv("PORT"))
}
