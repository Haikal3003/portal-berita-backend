package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        int       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`

	Articles []Article `json:"articles" gorm:"foreignKey:CategoryID"`
}

type ArticleCategory struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ArticleID  uuid.UUID `json:"article_id" gorm:"type:uuid"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`

	Article  Article  `json:"article" gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
