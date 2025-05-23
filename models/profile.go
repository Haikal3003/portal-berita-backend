package models

import (
	"time"
)

type Profile struct {
	ID            string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID        string    `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	Fullname      string    `json:"fullname" gorm:"not null"`
	Username      string    `json:"username" gorm:"not null;uniqueIndex"`
	Bio           string    `json:"bio"`
	Image         string    `json:"image"`
	ImagePublicID string    `json:"image_public_id"`
	Address       string    `json:"address"`
	BirthDate     string    `json:"birth_date"`
	User          User      `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
