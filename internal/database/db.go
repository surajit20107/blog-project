package database

import (
	"fmt"
	"log"

	"github.com/surajit/blog-project/config"
	"github.com/surajit/blog-project/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(
		&models.Comment{},
		&models.Post{},
		&models.Reaction{},
		&models.Tag{},
		&models.User{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("Database connected successfully")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("⚠️ Failed to close DB: %v", err)
		return
	}
	sqlDB.Close()
	log.Println("🛑 Database connection closed")
}
