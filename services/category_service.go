package services

import (
	"portal-berita-backend/models"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) FindOrCreateCategory(name string) (models.Category, error) {
	var category models.Category
	if err := s.DB.Where("LOWER(name) = LOWER(?)", name).First(&category).Error; err != nil {
		category = models.Category{
			Name: name,
			Slug: slug.Make(name),
		}
		if err := s.DB.Create(&category).Error; err != nil {
			return category, err
		}
	}
	return category, nil
}
