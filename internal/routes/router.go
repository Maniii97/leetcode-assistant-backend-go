package routes

import (
	"fmt"
	"lc-assistant-backend/internal/handlers"
	"lc-assistant-backend/internal/middleware"
	"lc-assistant-backend/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	// Initialize services
	genaiService, err := service.NewGeminiService(nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Gemini service: %v", err))
	}

	// Initialize handlers with dependencies
	chatHandler := handlers.NewChatHandler(genaiService)
	hintsHandler := handlers.NewHintsHandler(genaiService)
	reviewHandler := handlers.NewReviewHandler(genaiService)
	solutionHandler := handlers.NewSolutionHandler(genaiService)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		api.POST("/chat", chatHandler.GetAnswer)
		api.POST("/hints", hintsHandler.GetHints)
		api.POST("/codereview", reviewHandler.ReviewCode)
		api.POST("/solution", solutionHandler.AnalyseSolution)
	}

	return router
}
