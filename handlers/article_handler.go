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

type ArticleRequest struct {
	Title     string          `json:"title" validate:"required"`
	Content   string          `json:"content" validate:"required"`
	Thumbnail string          `json:"thumbnail"`
	Category  models.Category `json:"category"`
	Tags      []models.Tag    `json:"tags"`
}

// GET ALL ARTICLE
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

// GET ARTICLE BY ID
func (h *ArticleHandler) GetArticleByID(c *fiber.Ctx) error {
	articleID := c.Params("id")

	if err := h.ArticleService.IncrementArticleView(articleID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to increment view",
			"error":   err.Error(),
		})
	}

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

// CREATE ARTICLE
func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {

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

// UPDATE ARTICLE
func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
	articleID := c.Params("id")
	if articleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Article ID is required",
		})
	}

	var req ArticleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	article := &models.Article{}
	if err := h.CategoryService.DB.Preload("Tags").Preload("Category").Where("id = ?", articleID).First(&article).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Article not found",
		})
	}

	article.Title = req.Title
	article.Slug = slug.Make(req.Title)
	article.Content = req.Content
	article.Thumbnail = req.Thumbnail

	if req.Category.Name != "" {
		category, err := h.CategoryService.FindOrCreateCategory(req.Category.Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update category",
				"error":   err.Error(),
			})
		}
		article.CategoryID = category.ID
	}

	if len(req.Tags) > 0 {
		tags, err := h.TagService.FindOrCreateTags(req.Tags)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update tags",
				"error":   err.Error(),
			})
		}
		if err := h.TagService.DB.Model(&article).Association("Tags").Replace(tags); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to replace article tags",
				"error":   err.Error(),
			})
		}
	}

	if err := h.ArticleService.DB.Save(article).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update article",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article updated successfully",
		"data":    article,
	})
}

// DELETE
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

// GET ARTICLE BY CATEGORY
func (h *ArticleHandler) GetArticlesByCategory(c *fiber.Ctx) error {
	categoryName := c.Params("name")

	articles, err := h.ArticleService.FindArticlesByCategory(categoryName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Articles with category " + categoryName + " not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": articles,
	})

}

// GET ARTICLE BY TAG
func (h *ArticleHandler) GetArticlesByTag(c *fiber.Ctx) error {
	tagName := c.Params("name")

	articles, err := h.ArticleService.FindArticlesByTag(tagName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Article with tag " + tagName + " not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": articles,
	})
}

// PUBLISH ARTICLE
func (h *ArticleHandler) PublishArticle(c *fiber.Ctx) error {
	articleID := c.Params("id")

	if err := h.ArticleService.PublishArticle(articleID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Article not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article published successfully",
	})

}

// SAVE ARTICLE FOR USER
func (h *ArticleHandler) SaveArticle(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	articleID := c.Params("id")
	userRole := c.Locals("role")

	role := models.RoleType(userRole.(string))

	if err := h.ArticleService.SaveArticle(userID, articleID, role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article saved successfully",
	})
}

// GET SAVED ARTICLES
func (h *ArticleHandler) GetSavedArticles(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	savedArticles, err := h.ArticleService.GetSavedArticle(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get saved articles",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": savedArticles,
	})
}
