package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleType sebagai ENUM
type RoleType string

const (
	RoleAdmin RoleType = "ADMIN"
	RoleUser  RoleType = "USER"
)

// Model User
type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Role     RoleType  `json:"role" gorm:"type:varchar(20);not null;default:'USER'"`
}

type Profile struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Fullname   string    `json:"fullname" gorm:"not null"`
	Username   string    `json:"username" gorm:"not null"`
	Bio        string    `json:"bio"`
	Avatar_URL string    `json:"avatar_url"`
	Address    string    `json:"address"`
	Birthdate  string    `json:"birth_date"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_At time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

// BeforeCreate untuk Generate UUID otomatis
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
