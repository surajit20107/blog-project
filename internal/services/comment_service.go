package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/surajit/blog-project/internal/models"
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

func (r *CommentService) Add(postId string, userId *string, parentId *string, content string) (*models.Comment, error) {
	var post models.Post
	if err := r.db.First(&post, "id=?", postId).Error; err != nil {
		return nil, err
	}
	// Proceed with the comment creation
	now := time.Now()
	comment := &models.Comment{
		ID: uuid.New(),
		PostID: uuid.MustParse(postId),
		Content: content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if userId != nil {
		uid := uuid.MustParse(*userId)
		comment.UserID = &uid
	}
	if parentId != nil && *parentId != "" {
		var parent models.Comment
		if err := r.db.First(&parent, "id=?", *parentId).Error; err != nil  {
			return nil, err
		}
		pid := uuid.MustParse(*parentId)
		comment.ParentID = &pid
	}
	if err := r.repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (r CommentService) ListByPost(postId string) ([]models.Comment, error) {
	return r.repo.ListByPost(postId)
}
