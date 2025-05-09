package services

import (
	"portal-berita-backend/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Preload("profile").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetRoleAdmin() (*models.User, error) {
	admin := &models.User{}
	if err := s.DB.Preload("profile").Where("role = ?", "ADMIN").Find(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *UserService) GetRoleUser() (*models.User, error) {
	user := &models.User{}
	if err := s.DB.Preload("profile").Where("role = ?", "USER").Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
