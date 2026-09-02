package handler

import "github.com/surajit/blog-project/internal/services"

type ReactionHandler struct {
	Service *services.ReactionService
}

func NewReactionHandler(s *services.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		Service: s,
	}
}
