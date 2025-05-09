package seed

import (
	"log"
	"os"
	"portal-berita-backend/database"
	"portal-berita-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func SetupAdmin() {
	var count int64
	err := database.DB.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count).Error
	if err != nil {
		panic(err)
	}

	if count > 0 {
		log.Println("Admin user already exists.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
		return
	}

	adminUser := models.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: string(hashedPassword),
		Role:     models.RoleAdmin,
		Profile: &models.Profile{
			Fullname: os.Getenv("ADMIN_FULLNAME"),
			Username: os.Getenv("ADMIN_USERNAME"),
		},
	}

	if err := database.DB.Create(&adminUser).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Println("Admin user created successfully.")

}
