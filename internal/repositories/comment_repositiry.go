package repositories

import (
	"github.com/surajit/blog-project/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct{
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r CommentRepository) ListByPost(postId string) ([]models.Comment, error) {
	var cs []models.Comment
	if err := r.db.Where("post_id=?", postId).Find(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}
