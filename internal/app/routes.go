package app

import (
	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/config"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {}