package handlers

import (
	"portal-berita-backend/models"
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type ArticleHandler struct {
	ArticleService  *services.ArticleService
	CategoryService *services.CategoryService
	TagService      *services.TagService
}

func NewArticleHandler(articleService *services.ArticleService, categoryService *services.CategoryService, tagService *services.TagService) *ArticleHandler {
	return &ArticleHandler{
		ArticleService:  articleService,
		CategoryService: categoryService,
		TagService:      tagService,
	}
}

func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {
	users, err := h.ArticleService.GetArticles()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get article data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})

}

func (h *ArticleHandler) GetArticleByID(c *fiber.Ctx) error {
	articleID := c.Params("id")
	article, err := h.ArticleService.GetArticleByID(articleID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "article with ID: " + articleID + " not found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": article,
	})
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	type ArticleRequest struct {
		Title     string          `json:"title" validate:"required"`
		Content   string          `json:"content" validate:"required"`
		Thumbnail string          `json:"thumbnail"`
		Category  models.Category `json:"category"`
		Tags      []models.Tag    `json:"tags"`
	}

	var req ArticleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Input",
			"error":   err.Error(),
		})
	}

	authorID := c.Locals("userID")
	if authorID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	category, err := h.CategoryService.FindOrCreateCategory(req.Category.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process category",
			"error":   err.Error(),
		})
	}

	tags, err := h.TagService.FindOrCreateTags(req.Tags)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process tags",
			"error":   err.Error(),
		})
	}

	article := models.Article{
		Title:      req.Title,
		Slug:       slug.Make(req.Title),
		Content:    req.Content,
		Thumbnail:  req.Thumbnail,
		CategoryID: category.ID,
		AuthorID:   authorID.(string),
		Tags:       tags,
	}

	createdArticle, err := h.ArticleService.CreateArticle(&article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create article",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Article created successfully",
		"data":    createdArticle,
	})

}

func (h *ArticleHandler) DeleteArticleByID(c *fiber.Ctx) error {
	articleID := c.Params("id")

	if err := h.ArticleService.DeleteArticle(articleID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to delete article",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete article with ID: " + articleID + " successfully !",
	})
}
