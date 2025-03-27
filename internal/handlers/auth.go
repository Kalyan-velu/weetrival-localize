package handlers

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kalyan-velu/weetrival-localize/dto"
	"github.com/kalyan-velu/weetrival-localize/internal/auth"
)

// RegisterUser
// @BasePath /api/v1
// @Summary Register a new user
// @Description Creates a new user account after checking if the user already exists
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var req dto.CreateUserRequest
	ctx := context.Background()

	// Log raw request body
	buf, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf)) // Restore body for parsing

	if len(buf) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing request body"})
		return
	}

	// Bind JSON request to user model
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": "Missing or invalid request body"})
		return
	}

	log.Printf("Parsed request: %+v", req) // Now logs correctly

	// Register user
	registeredUser, err := auth.RegisterUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed", "message": err.Error()})
		return
	}

	// Successful response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    registeredUser,
	})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var req dto.LoginRequest
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}

	token, err := auth.LoginUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	auth.StoreTokenInCookie(c, token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ProtectedEndpoint is an example of a route that needs authentication
func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected endpoint!"})
}
