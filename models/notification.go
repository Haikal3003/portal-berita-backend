package models

import (
	"time"
)

type Notification struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"type:uuid;not null"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
