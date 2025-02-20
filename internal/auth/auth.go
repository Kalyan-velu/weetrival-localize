package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your-secret-key")

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

func RegisterUser(email, password string) (string, error) {
	return "", nil
}

// StoreTokenInCookie Store token in cookie and send response
func StoreTokenInCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
}
