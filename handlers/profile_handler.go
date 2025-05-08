package handlers

// import (
// 	"portal-berita-backend/services"

// 	"github.com/gofiber/fiber/v2"
// )

// type ProfileHandler struct {
// 	ProfileService *services.ProfileService
// }

// func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
// 	return &ProfileHandler{
// 		ProfileService: profileService,
// 	}
// }

// func (h *ProfileHandler) GetMyProfile(c *fiber.Ctx) error {
// 	userID := c.Locals("userID").(string)

// 	profile, err := h.ProfileService.GetProfileByUserID(userID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to get profile",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"profile": profile,
// 	})
// }
