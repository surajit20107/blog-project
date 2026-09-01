package services

import (
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
