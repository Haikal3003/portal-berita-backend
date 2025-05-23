package handlers

import (
	"portal-berita-backend/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService       *services.UserService
	CloudinaryService *services.CloudinaryService
}

func NewUserHandler(userService *services.UserService, cloudinaryService *services.CloudinaryService) *UserHandler {
	return &UserHandler{
		UserService:       userService,
		CloudinaryService: cloudinaryService,
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

	fullname := c.FormValue("fullname")
	username := c.FormValue("username")
	bio := c.FormValue("bio")
	address := c.FormValue("address")
	birthDate := c.FormValue("birth_date")

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	if fullname != "" {
		user.Profile.Fullname = fullname
	}

	if username != "" {
		user.Profile.Username = username
	}
	if bio != "" {
		user.Profile.Bio = bio
	}
	if address != "" {
		user.Profile.Address = address
	}
	if birthDate != "" {
		user.Profile.BirthDate = birthDate
	}

	fileHeader, err := c.FormFile("image")
	if err == nil && fileHeader != nil {
		if user.Profile.ImagePublicID != "" {
			_ = h.CloudinaryService.DeleteImage(user.Profile.ImagePublicID)
		}

		imageURL, publicID, err := h.CloudinaryService.UploadImage(fileHeader, "profile_image")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload image",
				"error":   err.Error(),
			})
		}
		user.Profile.Image = imageURL
		user.Profile.ImagePublicID = publicID
	}

	updatedProfile, err := h.UserService.UpdateUserProfile(userID, user.Profile)
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
