package services

import (
	"errors"
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

func (s *ArticleService) UpdateArticleById(articleID string, updatedArticle *models.Article) (*models.Article, error) {
	article := &models.Article{}

	if err := s.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
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

func (s *ArticleService) FindArticlesByCategory(name string) (*models.Category, error) {
	category := &models.Category{}

	if err := s.DB.Preload("Articles").Preload("Articles.Author.Profile").Where("LOWER(name) = LOWER(?)", name).First(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (s *ArticleService) FindArticlesByTag(name string) ([]models.Article, error) {
	var tag models.Tag

	err := s.DB.Preload("Articles.Tags").
		Preload("Articles.Author").
		Preload("Articles.Category").
		Where("name = ?", name).
		First(&tag).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	return tag.Articles, nil
}

func (s *ArticleService) PublishArticle(articleID string) error {
	article := &models.Article{}

	if err := s.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		return err
	}

	article.Status = models.StatusPublished

	if err := s.DB.Save(&article).Error; err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) GetSavedArticle(userID string) ([]models.SavedArticle, error) {
	var savedArticles []models.SavedArticle

	if err := s.DB.
		Where("user_id = ?", userID).
		Preload("Article.Author").
		Preload("Article.Category").
		Preload("Article.Tags").
		Find(&savedArticles).Error; err != nil {
		return nil, err
	}

	return savedArticles, nil
}

func (s *ArticleService) SaveArticle(userID, articleID string, role models.RoleType) error {
	if role != models.RoleUser {
		return errors.New("Only user can save article")
	}

	var article models.Article
	if err := s.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		return errors.New("Article not found")
	}

	var existing models.SavedArticle
	if err := s.DB.Where("user_id = ? AND article_id = ?", userID, articleID).First(&existing).Error; err != nil {
		return errors.New("Article already saved")
	}

	savedArticle := models.SavedArticle{
		UserID:    userID,
		ArticleID: articleID,
	}

	if err := s.DB.Create(&savedArticle).Error; err != nil {
		return errors.New("Failed to save article")
	}

	return nil

}

func (s *ArticleService) IncrementArticleView(articleID string) error {
	article := &models.Article{}

	if err := s.DB.Where("id : ?", articleID).First(&article).Error; err != nil {
		return err
	}

	article.Views += 1

	if err := s.DB.Save(&article).Error; err != nil {
		return err
	}

	return nil
}
