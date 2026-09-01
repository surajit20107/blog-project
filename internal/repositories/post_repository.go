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

func (r *PostRepository) GetAll() ([]*models.Post, error) {
	var posts []*models.Post
	if err := r.db.Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetById(id string) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("id=?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
