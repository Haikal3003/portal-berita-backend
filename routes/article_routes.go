package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(router fiber.Router, articleHandler *handlers.ArticleHandler) {
	articleRoute := router.Group("/articles", middlewares.JwtMiddleware())

	articleRoute.Get("/", articleHandler.GetAllArticles)
	articleRoute.Get("/:id", articleHandler.GetArticleByID)
	articleRoute.Get("/category/:name", articleHandler.GetArticlesByCategory)
	articleRoute.Get("/tag/:name", articleHandler.GetArticlesByTag)
	articleRoute.Post("/", articleHandler.CreateArticle)
	articleRoute.Put("/:id/publish", articleHandler.PublishArticle)
}
