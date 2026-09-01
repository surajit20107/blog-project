package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/surajit/blog-project/internal/auth"
	"github.com/surajit/blog-project/internal/models"
	"github.com/surajit/blog-project/internal/repositories"
)

type UserService struct{
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Create(username, email, password string) (*models.User, error) {
	u := models.User{
		ID: uuid.New(),
		UserName: username,
		Email: email,
		Password: password,
		DisplayName: username,
		Role: "reader",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	hashedPassword, err := auth.HashedPassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	if err := s.repo.Create(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
