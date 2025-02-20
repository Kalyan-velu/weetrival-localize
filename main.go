package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kalyan-velu/weetrival-localize/internal/handlers"
	"github.com/kalyan-velu/weetrival-localize/internal/middleware"
)

func main() {
	// Get PORT from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("⚠️ PORT environment variable not set. Using default:", port)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/register", handlers.RegisterUser)
		v1.POST("/login", handlers.LoginUser)

		// Middleware should be applied before handlers
		v1.Use(middleware.ABAC) // Apply ABAC middleware to all routes in v1

		v1.GET("/protected", handlers.ProtectedEndpoint)
	}

	// Start server on the correct port
	addr := fmt.Sprintf(":%s", port)
	log.Println("🚀 Server running on", addr)
	err := r.Run(addr)
	if err != nil {
		return
	}
}
