package handler

import "github.com/surajit/blog-project/internal/services"

type AuthHandler struct {
	services *services.UserService
}

func NewAuthHandler(services *services.UserService) *AuthHandler {
	return &AuthHandler{
		services: services,
	}
}

