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
		slugStr := slug.Make(inputTag.Name)
		var tag models.Tag

		err := s.DB.Where("slug = ?", slugStr).First(&tag).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				tag = models.Tag{
					Name: inputTag.Name,
					Slug: slugStr,
				}
				if err := s.DB.Create(&tag).Error; err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		tags = append(tags, tag)
	}
	return tags, nil
}
