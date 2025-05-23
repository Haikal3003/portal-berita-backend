package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(router fiber.Router, articleHandler *handlers.ArticleHandler) {
	articleRoute := router.Group("/articles", middlewares.JwtMiddleware())

	articleRoute.Get("/", articleHandler.GetAllArticles)
	articleRoute.Get("/saved-articles", articleHandler.GetSavedArticles)
	articleRoute.Get("/:id", articleHandler.GetArticleByID)
	articleRoute.Get("/category/:name", articleHandler.GetArticlesByCategory)
	articleRoute.Get("/tag/:name", articleHandler.GetArticlesByTag)

	articleRoute.Post("/", articleHandler.CreateArticle)
	articleRoute.Post("/:id/save", articleHandler.SaveArticle)

	articleRoute.Put("/:id/publish", articleHandler.PublishArticle)
	articleRoute.Put("/:id/update-article", articleHandler.UpdateArticle)

}
