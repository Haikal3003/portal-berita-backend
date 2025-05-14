package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler) {
	userRoute := router.Group("/users", middlewares.JwtMiddleware())

	userRoute.Get("/", userHandler.GetAllUsers)
	userRoute.Get("/:id", userHandler.GetUserById)
	userRoute.Put("/update-profile", userHandler.UpdateProfile)
}
