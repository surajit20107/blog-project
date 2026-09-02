package handler

import "github.com/surajit/blog-project/internal/services"

type AuthHandler struct {
	services *services.UserService
}

func NewAuthHandler(s *services.UserService) *AuthHandler {
	return &AuthHandler{
		services: s,
	}
}

