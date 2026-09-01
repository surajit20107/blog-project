package repositories

import "gorm.io/gorm"

type ReactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) *ReactionRepository {
	return &ReactionRepository{
		db: db,
	}
}
