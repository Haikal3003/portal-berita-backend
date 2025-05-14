package models

import (
	"time"
)

type Tag struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"uniqueIndex;not null"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;not null"`
	Articles  []Article `json:"articles" gorm:"many2many:article_tags"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
