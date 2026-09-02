package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/internal/services"
	"github.com/surajit/blog-project/internal/utils"
)

type PostHandler struct {
	Service *services.PostService
}

func NewPostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{
		Service: s,
	}
}

type CreatePostReq struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	var req CreatePostReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "Invalid Request")
	}
	userIDSrt, ok := c.Get("user_id").(string)
	if !ok {
		return utils.Err(c, http.StatusUnauthorized, "Invalid User ID") 
	}
	authorId, err := uuid.Parse(userIDSrt)
	if err != nil {
		return utils.Err(c, http.StatusUnauthorized, "Invalid User ID")
	}
	post, err := h.Service.Create(authorId, req.Title, req.Content, req.Tags)
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusCreated, true, "Post Created Successfully!", post)
}

func (h *PostHandler) GetAll(c echo.Context) error {
	posts, err := h.Service.GetAll()
	if err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "posts", posts)
}

func (h *PostHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		return utils.Err(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSON(c, http.StatusOK, true, "Post Deleted!", nil)
}
