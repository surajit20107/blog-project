package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/internal/services"
	"github.com/surajit/blog-project/internal/utils"
)

type AuthHandler struct {
	services *services.UserService
}

func NewAuthHandler(s *services.UserService) *AuthHandler {
	return &AuthHandler{
		services: s,
	}
}

type SignupReq struct {
	UserName    string `json:"user_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "Invalid Payload.")
	}
	u, err := h.services.Register(req.UserName, req.Email, req.Password)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, err.Error())
	}
	return utils.JSON(c, http.StatusCreated, true, "User Created Successfully!", u)
}

type LoginReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return utils.Err(c, http.StatusBadRequest, "Invalid Payload")
	}
	u, token, err := h.services.Login(req.Email, req.Password)
	if err != nil {
		return utils.Err(c, http.StatusBadRequest, "Invalid Credentials!")
	}
	return utils.JSON(c, http.StatusOK, true, "Login Successfully!", map[string]interface{}{
		"user":  u,
		"token": token,
	})
}
