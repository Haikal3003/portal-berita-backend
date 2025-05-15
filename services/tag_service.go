package services

import (
	"portal-berita-backend/models"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type TagService struct {
	DB *gorm.DB
}

func NewTagService(db *gorm.DB) *TagService {
	return &TagService{DB: db}
}

func (s *TagService) FindOrCreateTags(tagInputs []models.Tag) ([]models.Tag, error) {
	var tags []models.Tag
	for _, inputTag := range tagInputs {
		var tag models.Tag
		if err := s.DB.Where("name = ?", inputTag.Name).First(&tag).Error; err != nil {
			tag = models.Tag{
				Name: inputTag.Name,
				Slug: slug.Make(inputTag.Name),
			}
			if err := s.DB.Create(&tag).Error; err != nil {
				return nil, err
			}
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
