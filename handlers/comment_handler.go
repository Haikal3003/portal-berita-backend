package handlers

import (
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	CommentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		CommentService: commentService,
	}
}

func (h *CommentHandler) GetArticleComments(c *fiber.Ctx) error {
	articleID := c.Params("id")

	comments, err := h.CommentService.GetCommentsByArticleID(articleID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get comments",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": comments,
	})
}

func (h *CommentHandler) AddComment(c *fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userID").(string)

	type CommentRequest struct {
		Message string `json:"message"`
	}

	var req CommentRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Message is required",
		})
	}

	comment, err := h.CommentService.CreateComment(articleID, userID, req.Message)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create comment",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    comment,
		"message": "Add comment successfully !",
	})

}

func (h *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	commentID := c.Params("id")
	userID := c.Locals("userID").(string)

	if err := h.CommentService.DeleteComment(commentID, userID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to delete comment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete comment successfully !",
	})
}
