// internal/handlers/auth.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kalyan-velu/weetrival-localize/internal/auth"
	"net/http"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var creds auth.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Process registration
	token, err := auth.RegisterUser(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var creds auth.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := auth.LoginUser(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ProtectedEndpoint is an example of a route that needs authentication
func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected endpoint!"})
}
