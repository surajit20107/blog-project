package services

import "github.com/surajit/blog-project/internal/repositories"

type UserService struct{
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}