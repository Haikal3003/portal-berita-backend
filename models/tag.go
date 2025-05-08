package models

import (
	"time"
)

type Tag struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"uniqueIndex;not null"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Articles []Article `json:"articles" gorm:"many2many:article_tags"`
}

type ArticleTag struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleID int       `json:"article_id"`
	TagID     int       `json:"tag_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	Article Article `json:"article" gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag     Tag     `json:"tag" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
