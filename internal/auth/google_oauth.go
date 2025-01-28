// internal/auth/google_oauth.go
package auth

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"email"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin() {
	// Generate URL for Google login
	url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	log.Println("Visit the URL for Google login: ", url)
}
