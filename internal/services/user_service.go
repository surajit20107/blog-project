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

func (s *UserService) Register(username, email, password string) (*models.User, error) {
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

func (s *UserService) Login(email, password string) (*models.User, string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if !auth.CheckPassword(user.Password, password) {
		return nil, "", err
	}
	token, err := auth.GenerateJWT(user.ID.String(), user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
