package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin user moderator"`
}

// Load the secret key from environment variables
var secret = os.Getenv("AUTH_SECRET")

// Validate the secret and throw an error if missing
func init() {
	if secret == "" {
		panic("❌ AUTH_SECRET is not set in the environment variables!")
	}
}

var jwtSecret = []byte(secret)

// Credentials struct for user login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GenerateToken creates a JWT for a user
func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// LoginUser authenticates user and generates token
func LoginUser(email, password string) (string, error) {
	// Authenticate the user here (hash password, check DB)
	// For simplicity, we'll assume valid credentials here

	token, err := GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// RegisterUser function (Stub for now)
func RegisterUser(c *gin.Context) (string, error) {
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required info.", "message": err.Error()})
		return "", err
	}
	return "", nil

}

// StoreTokenInCookie stores token in cookie and sends response
func StoreTokenInCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
}
