package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null; uniqueIndex"`
	Fullname  string     `json:"fullname" gorm:"not null"`
	Username  string     `json:"username" gorm:"not null; uniqueIndex"`
	Bio       string     `json:"bio"`
	Image     string     `json:"image"`
	Address   string     `json:"address"`
	BirthDate *time.Time `json:"birth_date"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`

	User *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
