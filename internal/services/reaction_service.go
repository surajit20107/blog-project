package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/surajit/blog-project/internal/models"
	"github.com/surajit/blog-project/internal/repositories"
	"gorm.io/gorm"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
	db *gorm.DB
}

func NewReactionService(r *repositories.ReactionRepository, db *gorm.DB) *ReactionService {
	return &ReactionService {
		repo: r,
		db: db,
	}
}

func (s *ReactionService) TogglePostLike(userId string, postId string) (bool, error) {
	if re, err := s.repo.Find(userId, postId); err == nil && re != nil {
		if err := s.repo.Delete(re.Id.String()); err != nil {
			return false, err
		}
		return false, nil
	}
	nr := &models.Reaction{
		Id: uuid.New(),
		UserID: uuid.MustParse(userId),
		PostID: uuid.MustParse(postId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.Create(nr); err != nil {
		return false, nil
	}
	return true, nil
}
