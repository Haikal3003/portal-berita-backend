package services

import (
	"portal-berita-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) CreateUser(fullname, username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
		Profile: models.Profile{
			Fullname: fullname,
			Username: username,
		},
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user := &models.User{}
	if err := s.DB.Preload("Profile").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
