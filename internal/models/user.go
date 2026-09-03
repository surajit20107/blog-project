package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserName    string    `gorm:"unique;not null" json:"username"`
	Email       string    `gorm:"unique;not null" json:"email"`
	Password    string    `gorm:"not null"      json:"-"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	Role        string    `gorm:"default:reader" json:"role"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Posts []Post `gorm:"foreignKey:AuthorID" json:"posts,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments"`
	Reacts []Reaction `gorm:"foreignKey:UserID" json:"reactions"`
}
