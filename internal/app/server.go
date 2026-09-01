package app

import (
	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/config"
	"gorm.io/gorm"
)

type Server struct {
	E *echo.Echo
	DB *gorm.DB
	Config *config.Config
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	e := echo.New()
	RegisterRoutes(e, db, cfg)
	return &Server{
		E: e,
		DB: db,
		Config: cfg,
	}
}

func (s *Server) Start(addr string) error {
	return s.E.Start(addr)
}