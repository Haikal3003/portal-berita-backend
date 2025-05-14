package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")

	auth.Post("/register", authHandler.RegisterUser)
	auth.Post("/login", authHandler.LoginUser)
	auth.Post("/change-password", middlewares.JwtMiddleware(), authHandler.ChangePassword)
}
