package handler

import "github.com/surajit/blog-project/internal/services"

type CommentHandler struct {
	Service *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{
		Service: s,
	}
}