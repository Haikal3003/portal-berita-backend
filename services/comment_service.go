package services

import (
	"portal-berita-backend/models"

	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		DB: db,
	}
}

func (s *CommentService) CreateComment(articleID, userID, message string) (*models.Comment, error) {
	comment := &models.Comment{
		Message:   message,
		UserID:    userID,
		ArticleID: articleID,
	}

	if err := s.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Preload("User.Profile").Where("id = ?", comment.ID).First(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil

}

func (s *CommentService) DeleteComment(commentID, userID string) error {
	var comment models.Comment

	if err := s.DB.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
		return err
	}

	return s.DB.Delete(&comment).Error

}

func (s *ArticleService) GetCommentsByArticleID(articleID string) ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.DB.Preload("User.Profile").Where("article_id = ?", articleID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
