package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/internal/services"
	"github.com/surajit/blog-project/internal/utils"
)

type ReactionHandler struct {
	Service *services.ReactionService
}

func NewReactionHandler(s *services.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		Service: s,
	}
}

type ReactionReq struct {
	Type string `json:"type"`
}

func (h *ReactionHandler) Toggle(c echo.Context) error {
	postID := c.Param("id")
	var req ReactionReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "Invalid Payload!")
	}
	uid := c.Get("user_id")
	if uid == nil {
		return utils.Err(c, http.StatusUnauthorized, "Unauthorized!")
	}
	userID := uid.(string)
	liked, err := h.Service.TogglePostLike(userID, postID)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	if liked {
		return utils.JSON(c, http.StatusOK, true, "Post Liked!", nil)
	}
	return utils.JSON(c, http.StatusOK, true, "Like Removed!", nil)
}
