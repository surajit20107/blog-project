package utils

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c echo.Context, data interface{}, msg string) error {
	if msg == "" {
		msg = "OK"
	}
	return c.JSON(http.StatusOK, APIResponse{Success: true, Message: msg, Data: data})
}

func Error(c echo.Context, code int, msg string) error {
	return c.JSON(code, APIResponse{Success: false, Message: msg})
}

func ParseIntQuery(c echo.Context, key string, def int) int {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	if n, err := strconv.Atoi(v); err == nil {
		return n
	}
	return def
}

func JSON(c echo.Context, code int, success bool, msg string, data interface{}) error {
	return c.JSON(code, APIResponse{
		Success: success,
		Message: msg,
		Data:    data,
	})
}

func OK(c echo.Context, data interface{}) error {
	return JSON(c, http.StatusOK, true, "OK", data)
}

func Err(c echo.Context, code int, msg string) error {
	return JSON(c, code, false, msg, nil)
}

func MakeSlug(s string) string {
	return GenerateSlug(s)
}
