package services

import (
	"portal-berita-backend/models"

	"github.com/gosimple/slug"

	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{
		DB: db,
	}
}

func (s *ArticleService) GetArticles() ([]models.Article, error) {
	var articles []models.Article
	if err := s.DB.Preload("Author").Preload("Category").Preload("Tags").Preload("Comments").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *ArticleService) GetArticleByID(articleID string) (*models.Article, error) {
	article := &models.Article{}
	if err := s.DB.Preload("Author").Preload("Category").Preload("Tags").Preload("Comments").Where("id = ?", articleID).First(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) CreateArticle(article *models.Article) (*models.Article, error) {
	if err := s.DB.Create(article).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Preload("Author.Profile").
		Preload("Category").
		Preload("Tags").
		Preload("Comments").
		Where("id = ?", article.ID).
		First(article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) UpdateArticle(authorID string, updatedArticle *models.Article) (*models.Article, error) {
	article := &models.Article{}

	if err := s.DB.Where("author_id = ?", authorID).First(&article).Error; err != nil {
		return nil, err
	}

	updatedArticle.Title = article.Title
	updatedArticle.Slug = slug.Make(updatedArticle.Title)
	updatedArticle.Content = article.Content
	updatedArticle.Thumbnail = article.Thumbnail

	if err := s.DB.Save(article).Error; err != nil {
		return nil, err
	}

	return article, nil

}

func (s *ArticleService) DeleteArticle(articleID string) error {
	if err := s.DB.Delete(&models.Article{}, articleID).Error; err != nil {
		return err
	}

	return nil
}
