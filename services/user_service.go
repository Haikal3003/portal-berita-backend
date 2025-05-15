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
	if err := s.DB.Preload("Profile").
		Preload("Articles").
		Preload("Saved").
		Preload("Notifications").
		Preload("Comments").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	user := &models.User{}
	if err := s.DB.Preload("Profile").
		Preload("Articles").
		Preload("Saved").
		Preload("Notifications").
		Preload("Comments").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUsersByRole(role string) ([]models.User, error) {
	var users []models.User
	if err := s.DB.Preload("Profile").
		Preload("Articles").
		Preload("Saved").
		Preload("Notifications").
		Preload("Comments").Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUserProfile(userID string, updatedProfile *models.Profile) (*models.Profile, error) {
	profile := &models.Profile{}

	if err := s.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	profile.Fullname = updatedProfile.Fullname
	profile.Username = updatedProfile.Username
	profile.Bio = updatedProfile.Bio
	profile.Image = updatedProfile.Image
	profile.Address = updatedProfile.Address
	profile.BirthDate = updatedProfile.BirthDate

	if err := s.DB.Save(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}
