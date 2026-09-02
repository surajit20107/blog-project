package handler

import "github.com/surajit/blog-project/internal/services"

type PostHandler struct {
	Service *services.PostService
}

func NewPostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{
		Service: s,
	}
}
