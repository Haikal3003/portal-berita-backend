package models

import (
	"time"

	"github.com/google/uuid"
)

type RoleType string

const (
	RoleAdmin RoleType = "ADMIN"
	RoleUser  RoleType = "USER"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email" gorm:"uniqueIndex;not null"`
	Password string    `json:"password" gorm:"not null"`
	Role     RoleType  `json:"role" gorm:"type:VARCHAR(10);default:'USER'"`
	Profile  Profile   `json:"profile"`
	Articles []Article `json:"articles"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
