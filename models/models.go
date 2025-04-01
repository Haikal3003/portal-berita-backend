package models

import (
	"time"

	"github.com/google/uuid"
)

// Model User
type RoleType string

const (
	RoleAdmin RoleType = "ADMIN"
	RoleUser  RoleType = "USER"
)

type User struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email         string         `json:"email" gorm:"unique;not null"`
	Password      string         `json:"password" gorm:"not null"`
	Role          RoleType       `json:"role" gorm:"type:varchar(20);not null;default:'USER'"`
	Profile       Profile        `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Articles      []Article      `json:"articles" gorm:"foreignKey:AuthorID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
	Comments      []Comment      `json:"comments" gorm:"foreignKey:UserID"`
	Likes         []Like         `json:"likes" gorm:"foreignKey:UserID"`
	Bookmarks     []Bookmark     `json:"bookmarks" gorm:"foreignKey:UserID"`
}

// Model Profile
type Profile struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;unique;not null"`
	Fullname  string    `json:"fullname" gorm:"not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Address   string    `json:"address"`
	Birthdate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Article
type StatusType string

const (
	StatusDraft       StatusType = "DRAFT"
	StatusUnpublished StatusType = "UNPUBLISHED"
	StatusPublished   StatusType = "PUBLISHED"
	StatusArchived    StatusType = "ARCHIVED"
)

type Article struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string     `json:"title" gorm:"not null"`
	Slug      string     `json:"slug" gorm:"not null;unique"`
	Content   string     `json:"content" gorm:"not null"`
	Thumbnail string     `json:"thumbnail"`
	AuthorID  uuid.UUID  `json:"author_id" gorm:"type:uuid;not null"`
	Author    User       `json:"author" gorm:"foreignKey:AuthorID"`
	Status    StatusType `json:"status" gorm:"type:varchar(20);not null;default:'UNPUBLISHED'"`
	Views     int        `json:"views"`
	Comments  []Comment  `json:"comments" gorm:"foreignKey:ArticleID"`
	Likes     []Like     `json:"likes" gorm:"foreignKey:ArticleID"`
	Bookmarks []Bookmark `json:"bookmarks" gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Category
type Category struct {
	ID        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"not null;unique"`
	Articles  []Article `json:"articles" gorm:"many2many:article_categories;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Tag
type Tag struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"not null;unique"`
	Articles  []Article `json:"articles" gorm:"many2many:article_tags;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Notification
type Notification struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Comment
type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Model Like
type Like struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// Model Bookmark
type Bookmark struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
