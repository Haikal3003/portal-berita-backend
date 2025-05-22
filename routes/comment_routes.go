package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(router fiber.Router, commentHandler *handlers.CommentHandler) {
	commentRoutes := router.Group("comment", middlewares.JwtMiddleware())

	commentRoutes.Get("/:id", commentHandler.GetArticleComments)
	commentRoutes.Post("/add-comment/:id", commentHandler.AddComment)
	commentRoutes.Delete("/", commentHandler.DeleteComment)
}
