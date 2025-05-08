package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusType string

const (
	StatusDraft       StatusType = "DRAFT"
	StatusUnpublished StatusType = "UNPUBLISHED"
	StatusPublished   StatusType = "PUBLISHED"
	StatusArchived    StatusType = "ARCHIVED"
)

type Article struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title     string    `json:"title" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	Thumbnail string    `json:"thumbnail"`
	AuthorID  uuid.UUID `json:"author_id" gorm:"not null"`
	Status    string    `json:"status" gorm:"type:VARCHAR(20); default:'DRAFT'"`
	Views     int       `json:"views" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`

	Author User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CategoryID int      `json:"category_id" gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Tags []Tag `json:"tags" gorm:"many2many:article_tags"`
}

// for role user
type SavedArticle struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	User    User    `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Article Article `json:"article" gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
