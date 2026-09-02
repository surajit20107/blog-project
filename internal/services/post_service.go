package services

import (
	"github.com/google/uuid"
	"github.com/surajit/blog-project/internal/models"
	"github.com/surajit/blog-project/internal/repositories"
	"github.com/surajit/blog-project/internal/utils"
	"gorm.io/gorm"
	"time"
)

type PostService struct {
	repo *repositories.PostRepository
	db   *gorm.DB
}

func NewPostService(r *repositories.PostRepository, db *gorm.DB) *PostService {
	return &PostService{
		repo: r,
		db:   db,
	}
}

func (s *PostService) Create(authorId uuid.UUID, title, content string, tags []string) (*models.Post, error) {
	post := &models.Post{
		ID:       uuid.New(),
		AuthorID: authorId,
		Title:    title,
		Slug:     utils.GenerateSlug(title + "-" + uuid.NewString()[:6]),
		Status: "published",
		Visibility: "public",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		for _, tagName := range tags {
			var tag models.Tag
			if err := tx.Where("name=?", tagName).First(&tag).Error; err != nil {
				tag = models.Tag{
					ID:   uuid.New(),
					Name: tagName,
					Slug: utils.GenerateSlug(tagName),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
			}
			if err := tx.Model(post).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) Delete(postId string) error {
	return s.repo.Delete(postId)
}

func (s *PostService) Update(postId string, title, content string, tags []string) (*models.Post, error) {
	post, err := s.repo.GetById(postId)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Content = content
	post.Slug = utils.GenerateSlug(title + "-" + uuid.NewString()[:6])
	post.UpdatedAt = time.Now()

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err :=  tx.Save(&post).Error; err != nil {
			return err
		}
		if len(tags) > 0 {
			var newTags []models.Tag
			for _, tagName := range tags {
				var tag models.Tag
				if err := tx.Where("name=?", tagName).First(&tag).Error; err != nil {
					tag = models.Tag{
						ID: uuid.New(),
						Name: tagName,
						Slug: utils.GenerateSlug(tagName),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}
					if err := tx.Create(&tag).Error; err != nil {
						return err
					}
				}
				newTags =  append(newTags, tag)
			}
			if err := tx.Model(&post).Association("Tags").Clear(); err != nil {
				return err
			}
			if err := tx.Model(&post).Association("Tags").Append(newTags); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetAll() ([]*models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetById(postId string) (*models.Post, error) {
	return s.repo.GetById(postId)
}
