package handlers

import (
	"portal-berita-backend/services"
	"portal-berita-backend/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	type RegisterInput struct {
		Fullname string `json:"fullname" validate:"required"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	user, err := h.AuthService.CreateUser(input.Fullname, input.Username, input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    user,
		"message": "Registered successfully",
	})
}

func (h *AuthHandler) LoginUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	user, err := h.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	c.Set("Authorization", "Bearer "+token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    user,
		"token":   token,
		"message": "Login successful",
	})
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {

	type ChangePasswordInput struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	var input ChangePasswordInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userID := c.Locals("userID").(string)

	if err := h.AuthService.ChangePassword(userID, input.OldPassword, input.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Change password successfully!",
	})
}
