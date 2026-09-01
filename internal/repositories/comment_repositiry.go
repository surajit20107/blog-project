package repositories

import "gorm.io/gorm"

type CommentRepository struct{
	db *gorm.DB
}