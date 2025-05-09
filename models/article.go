package models

import (
	"time"
)

type StatusType string

const (
	StatusDraft       StatusType = "DRAFT"
	StatusUnpublished StatusType = "UNPUBLISHED"
	StatusPublished   StatusType = "PUBLISHED"
	StatusArchived    StatusType = "ARCHIVED"
)

type Article struct {
	ID         string     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title      string     `json:"title" gorm:"not null"`
	Slug       string     `json:"slug" gorm:"not null"`
	Content    string     `json:"content" gorm:"not null"`
	Thumbnail  string     `json:"thumbnail"`
	AuthorID   string     `json:"author_id" gorm:"type:uuid;not null"`
	Author     User       `json:"author" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Status     StatusType `json:"status" gorm:"type:VARCHAR(20);default:'DRAFT'"`
	Views      int        `json:"views" gorm:"default:0"`
	CategoryID int        `json:"category_id" gorm:"not null"`
	Category   Category   `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags       []Tag      `json:"tags" gorm:"many2many:article_tags"`
	Comments   []Comment  `json:"comments" gorm:"foreignKey:ArticleID"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
