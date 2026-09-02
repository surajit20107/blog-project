package repositories

import (
	"github.com/surajit/blog-project/internal/models"
	"gorm.io/gorm"
)

type ReactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) *ReactionRepository {
	return &ReactionRepository{
		db: db,
	}
}

func (r *ReactionRepository) Find(userId, postId string) (*models.Reaction, error) {
	var re models.Reaction
	if err := r.db.Where("user_id=? AND post_id=?", userId, postId).First(&re).Error; err != nil {
		return nil, err
	}
	return &re, nil
}

func (r *ReactionRepository) Create(re *models.Reaction) error {
	return r.db.Create(&re).Error
}

func (r *ReactionRepository) Delete(id string) error {
	return r.db.Delete(&models.Reaction{}, "id=?", id).Error
}
