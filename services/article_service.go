package services

import (
	"portal-berita-backend/models"
	"strings"

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
	err := s.DB.Preload("Author.Profile").Preload("Categories").Preload("Tags").Find(&articles).Error
	return articles, err
}

func (s *ArticleService) GetArticleByID(id string) (*models.Article, error) {
	var article models.Article
	err := s.DB.Preload("Author.Profile").Preload("Categories").Preload("Tags").First(&article, "id = ?", id).Error
	return &article, err
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	if err := s.DB.Create(article).Error; err != nil {
		return err
	}
	return s.DB.Preload("Author.Profile").First(article, article.ID).Error
}

func (s *ArticleService) UpdateArticle(id string, updated *models.Article) error {
	var article models.Article
	if err := s.DB.First(&article, "id = ?", id).Error; err != nil {
		return err
	}

	updated.ID = article.ID
	return s.DB.Model(&article).Updates(updated).Error
}

func (s *ArticleService) DeleteArticle(id string) error {
	return s.DB.Delete(&models.Article{}, "id = ?", id).Error
}

func (s *ArticleService) GetArticlesByCategory(categoryID string) ([]models.Article, error) {
	var articles []models.Article
	err := s.DB.Joins("JOIN article_categories ON article_categories.article_id = articles.id").
		Where("article_categories.category_id = ?", categoryID).
		Preload("Author.Profile").
		Preload("Categories").
		Preload("Tags").
		Find(&articles).Error
	return articles, err
}

func (s *ArticleService) GetArticlesByTag(tagID string) ([]models.Article, error) {
	var articles []models.Article
	err := s.DB.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID).
		Preload("Author.Profile").
		Preload("Categories").
		Preload("Tags").
		Find(&articles).Error
	return articles, err
}

func (s *ArticleService) SearchArticles(keyword string) ([]models.Article, error) {
	var articles []models.Article
	err := s.DB.Where("LOWER(title) ILIKE ?", "%"+strings.ToLower(keyword)+"%").
		Preload("Author.Profile").
		Preload("Categories").
		Preload("Tags").
		Find(&articles).Error
	return articles, err
}
