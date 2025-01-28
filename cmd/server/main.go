package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalyan-velu/weetrival-localize/internal/handlers"
	"github.com/kalyan-velu/weetrival-localize/internal/middleware"
)

func main() {
	r := gin.Default()

	// Apply ABAC middleware to specific routes
	r.Use(middleware.ABAC)

	// Set up routes
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/protected", handlers.ProtectedEndpoint)

	// Start server
	r.Run(":8080")
}
