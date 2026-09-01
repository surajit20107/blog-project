package repositories

import (
	"github.com/surajit/blog-project/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct{
	db *gorm.DB
	
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}
