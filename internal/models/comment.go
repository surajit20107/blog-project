package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PostID     uuid.UUID  `gorm:"type:uuid;not null" json:"post_id"`
	UserID     *uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ParentID   *uuid.UUID `gorm:"type:uuid" json:"parent_uuid"`
	Content    string     `gorm:"type:text" json:"content"`
	IsApproved bool       `gorm:"default:false" json:"is_approved"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
