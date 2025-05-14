package models

import (
	"time"
)

type RoleType string

const (
	RoleAdmin RoleType = "ADMIN"
	RoleUser  RoleType = "USER"
)

type User struct {
	ID            string         `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email         string         `json:"email" gorm:"uniqueIndex;not null"`
	Password      string         `json:"password" gorm:"not null"`
	Role          RoleType       `json:"role" gorm:"type:VARCHAR(10);default:'USER'"`
	Profile       *Profile       `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Articles      []Article      `json:"articles" gorm:"foreignKey:AuthorID"`
	Saved         []SavedArticle `json:"saved_articles" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
	Comments      []Comment      `json:"comments" gorm:"foreignKey:UserID"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}
