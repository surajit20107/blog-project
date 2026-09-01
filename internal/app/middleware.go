package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/config"
	"github.com/surajit/blog-project/internal/auth"
	"gorm.io/gorm"
)

func InjectDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Missing Authorization header",
				})
			}
			if !strings.HasPrefix(authHeader, "Bearer") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid Authorization header",
				})
			}
			token := strings.TrimPrefix(authHeader, "Bearer")
			claims, err := auth.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid Token",
				})
			}
			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)
			return next(c)
		}
	}
}
