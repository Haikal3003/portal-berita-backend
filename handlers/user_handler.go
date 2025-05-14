package handlers

import (
	"portal-berita-backend/models"
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserService.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to get user data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := h.UserService.GetUserByID(userID)

	if err != nil {
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"message": "User with ID: " + userID + "not found !",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	type UpdateProfileInput struct {
		Fullname  string `json:"fullname"`
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Address   string `json:"address"`
		BirthDate string `json:"birth_date"`
	}

	var input UpdateProfileInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	profile := &models.Profile{
		Fullname:  input.Fullname,
		Username:  input.Username,
		Bio:       input.Bio,
		Image:     input.Image,
		Address:   input.Address,
		BirthDate: input.BirthDate,
	}

	updatedProfile, err := h.UserService.UpdateUserProfile(userID, profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update profile",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile updated successfully",
		"data":    updatedProfile,
	})
}
