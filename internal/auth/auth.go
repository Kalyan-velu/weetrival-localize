package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/kalyan-velu/weetrival-localize/dto"
	"github.com/kalyan-velu/weetrival-localize/internal/models"
	"github.com/kalyan-velu/weetrival-localize/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

// Load environment variables & validate secret
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, relying on system environment")
	}

	secret := os.Getenv("AUTH_SECRET")
	if secret == "" {
		log.Fatal("❌ AUTH_SECRET is missing in environment variables")
	}
	jwtSecret = []byte(secret)
}

// Credentials struct for user login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GenerateToken creates a JWT for a user
func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub": email,
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// LoginUser authenticates user and generates token
func LoginUser(email, password string) (string, error) {
	// TODO: Validate user from DB (check hashed password)
	// Assume user is valid for now

	token, err := GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// RegisterUser validates and registers a new user
func RegisterUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {
	// Validate user input
	if len(req.Name) < 3 || len(req.Name) > 50 {
		return nil, errors.New("name must be between 3 and 50 characters")
	}
	if len(req.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	//Check if user already exists
	existingUser, err := repositories.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	hashedPasswordString := string(hashedPassword)

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: &hashedPasswordString,
		Role:         req.Role,
		CreatedAt:    time.Now(),
	}
	//** Validate role before saving
	if err := user.ValidateRole(); err != nil {
		return nil, err
	}

	// Save the user to the database
	if err := repositories.CreateUser(ctx, user); err != nil {
		return nil, errors.New("failed to save user to the database")
	}

	return user, nil
}

// StoreTokenInCookie stores JWT token in a cookie
func StoreTokenInCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", os.Getenv("COOKIE_DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
}
