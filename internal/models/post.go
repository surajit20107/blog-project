package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AuthorID      uuid.UUID  `gorm:"type:uuid;not null" json:"author_id"`
	Author        User       `gorm:"foreignKey:AuthorID" json:"author"`
	Title         string     `gorm:"not null" json:"title"`
	Slug          string     `gorm:"unique;not null" json:"slug"`
	Excerpt       string     `json:"excerpt"`
	Content       string     `gorm:"type:text" json:"content"`
	Status        string     `gorm:"default:draft" json:"status"`
	Visibility    string     `gorm:"default:public" json:"visibility"`
	PublishedAt   *time.Time `json:"published_at"`
	Views         int64      `gorm:"defult:0" json:"views"`
	CommentCount  int64      `gorm:"defult:0" json:"comment_count"`
	ReactionCount int64      `gorm:"defult:0" json:"reaction_count"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
