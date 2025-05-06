package routes

import (
	"portal-berita-backend/handlers"
	"portal-berita-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(router fiber.Router, articleHandler *handlers.ArticleHandler) {
	article := router.Group("/articles", middlewares.JwtMiddleware())

	article.Get("/", articleHandler.GetAllArticles)
	article.Get("/:id", articleHandler.GetArticleByID)
	article.Post("/", articleHandler.CreateArticle)
	article.Put("/:id", articleHandler.UpdateArticle)
	article.Delete("/:id", articleHandler.DeleteArticle)
	article.Get("/category/:categoryID", articleHandler.GetArticlesByCategory)
	article.Get("/tag/:tagID", articleHandler.GetArticlesByTag)
	article.Get("/search?query=:query", articleHandler.SearchArticles)
}
