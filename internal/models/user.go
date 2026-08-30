package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserName    string    `gorm:"unique;not null" json:"username"`
	Email       string    `gorm:"unique;not null" json:"email"`
	Password    string    `gorm:"not null"      json:"-"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	Role        string    `gorm:"default:reader" json:"role"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
