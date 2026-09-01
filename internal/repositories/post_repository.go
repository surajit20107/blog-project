package repositories

import "gorm.io/gorm"

type PostRepository struct{
	db *gorm.DB
	
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}
