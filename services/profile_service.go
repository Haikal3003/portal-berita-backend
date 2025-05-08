package services

// import (
// 	"portal-berita-backend/models"

// 	"gorm.io/gorm"
// )

// type ProfileService struct {
// 	DB *gorm.DB
// }

// func NewProfileService(db *gorm.DB) *ProfileService {
// 	return &ProfileService{
// 		DB: db,
// 	}
// }

// func (s *ProfileService) GetProfileByUserID(userID string) (*models.Profile, error) {
// 	profile := &models.Profile{}
// 	if err := s.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
// 		return nil, err
// 	}

// 	return profile, nil
// }
