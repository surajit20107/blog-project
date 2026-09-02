package app

import (
	"github.com/labstack/echo/v4"
	"github.com/surajit/blog-project/config"
	"github.com/surajit/blog-project/internal/handler"
	"github.com/surajit/blog-project/internal/models"
	"github.com/surajit/blog-project/internal/repositories"
	"github.com/surajit/blog-project/internal/services"
	"github.com/surajit/blog-project/internal/utils"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	e.Use(InjectDB(db))

	// User Repos or Services Call
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService)

	// Post Repos or Services Call
	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository, db)
	postHandler := handler.NewPostHandler(postService)

	// Comment Repos or Services Call
	commentRepository := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository, db)
	commentHandler := handler.NewCommentHandler(commentService)

	// Reaction Repos or Services Call
	reactionRepository := repositories.NewReactionRepository(db)
	reactionService := services.NewReactionService(reactionRepository, db)
	reactionHandler := handler.NewReactionHandler(reactionService)

	// Public Endpoints
	e.POST("/api/v1/signup", authHandler.SignUp)             // User Registration
	e.POST("/api/v1/login", authHandler.Login)               // User Login
	e.GET("/api/v1/posts", postHandler.GetAllPosts)          // Get All Posts
	e.GET("/api/v1/posts/:id", postHandler.GetPostById)      // Get Post By ID
	e.GET("/api/v1/posts/:id/comments", commentHandler.List) // Get Comments By Post ID
	e.PUT("/api/v1/posts/:id", postHandler.UpdatePost)       // Update Post

	// Get All Tags
	e.GET("/api/v1/tags", func(c echo.Context) error {
		db := c.Get("db").(*gorm.DB)
		var tags []models.Tag
		_ = db.Find(&tags).Error
		return utils.JSON(c, 200, true, "tags", tags)
	})

	// Protected Routes
	g := e.Group("/api/v1")
	g.Use(JWTMiddleware(cfg))
	g.POST("/posts", postHandler.CreatePost)                 // Create Post
	g.DELETE("/posts/:id", postHandler.Delete)               // Delete Post
	g.POST("/posts/:id/comments", commentHandler.AddComment) // Add Comment
	g.POST("/posts/:id/reactions", reactionHandler.Toggle)   // Toggle Reaction
}
