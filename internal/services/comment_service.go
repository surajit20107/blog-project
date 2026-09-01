package services

import (
	"github.com/surajit/blog-project/internal/repositories"
	"gorm.io/gorm"
)

type CommentService struct{
	repo *repositories.CommentRepository
	db *gorm.DB
}

func NewCommentService(r *repositories.CommentRepository, db *gorm.DB) *CommentService {
	return  &CommentService{
		repo: r,
		db: db,
	}
}