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

type StatusType string

const (
	StatusDraft       StatusType = "DRAFT"
	StatusUnpublished StatusType = "UNPUBLISHED"
	StatusPublished   StatusType = "PUBLISHED"
	StatusArchived    StatusType = "ARCHIVED"
)

// =======================
// User
// =======================
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

// =======================
// Profile
// =======================
type Profile struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;unique;not null"`
	User      *User     `json:"user" gorm:"foreignKey:UserID"`
	Fullname  string    `json:"fullname" gorm:"not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Address   string    `json:"address"`
	Birthdate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Article
// =======================
type Article struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title      string     `json:"title" gorm:"not null"`
	Slug       string     `json:"slug" gorm:"not null;unique"`
	Content    string     `json:"content" gorm:"not null"`
	Thumbnail  string     `json:"thumbnail"`
	AuthorID   uuid.UUID  `json:"author_id" gorm:"type:uuid;not null"`
	Author     User       `json:"author" gorm:"foreignKey:AuthorID"`
	Status     StatusType `json:"status" gorm:"type:varchar(20);not null;default:'UNPUBLISHED'"`
	Views      int        `json:"views"`
	Comments   []Comment  `json:"comments" gorm:"foreignKey:ArticleID"`
	Likes      []Like     `json:"likes" gorm:"foreignKey:ArticleID"`
	Bookmarks  []Bookmark `json:"bookmarks" gorm:"foreignKey:ArticleID"`
	Categories []Category `json:"categories" gorm:"many2many:article_categories"`
	Tags       []Tag      `json:"tags" gorm:"many2many:article_tags"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Category
// =======================
type Category struct {
	ID        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"not null;unique"`
	Articles  []Article `json:"articles" gorm:"many2many:article_categories"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Tag
// =======================
type Tag struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"not null;unique"`
	Articles  []Article `json:"articles" gorm:"many2many:article_tags"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Notification
// =======================
type Notification struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Comment
// =======================
type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// =======================
// Like
// =======================
type Like struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex:idx_user_article_like"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null;uniqueIndex:idx_user_article_like"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// =======================
// Bookmark
// =======================
type Bookmark struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex:idx_user_article_bookmark"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null;uniqueIndex:idx_user_article_bookmark"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
