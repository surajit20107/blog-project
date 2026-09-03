package main

import (
	"fmt"
	"log"

	"github.com/surajit/blog-project/config"
	"github.com/surajit/blog-project/internal/app"
	"github.com/surajit/blog-project/internal/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect with Database: %v", err)
	}
	defer database.CloseDB(db)

	// Initialize Echo server
	server := app.NewServer(db, cfg)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("🚀 Starting server on: %s\n", addr)

	if err := server.Start(addr); err != nil {
		log.Fatal(err)
	}
}
