package handlers

import (
	"portal-berita-backend/models"
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	ArticleService *services.ArticleService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		ArticleService: articleService,
	}
}

func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {
	articles, err := h.ArticleService.GetArticles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch articles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}

func (h *ArticleHandler) GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	article, err := h.ArticleService.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.ArticleService.CreateArticle(&article); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create article",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedArticle models.Article
	if err := c.BodyParser(&updatedArticle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.ArticleService.UpdateArticle(id, &updatedArticle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update article",
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedArticle)

}

func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.ArticleService.DeleteArticle(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete article",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Article deleted successfully",
	})
}

func (h *ArticleHandler) GetArticlesByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	articles, err := h.ArticleService.GetArticlesByCategory(categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch articles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}

func (h *ArticleHandler) GetArticlesByTag(c *fiber.Ctx) error {
	tagID := c.Params("tag_id")
	articles, err := h.ArticleService.GetArticlesByTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch articles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}

func (h *ArticleHandler) SearchArticles(c *fiber.Ctx) error {
	keyword := c.Query("query")

	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Keyword is required",
		})
	}

	articles, err := h.ArticleService.SearchArticles(keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}
