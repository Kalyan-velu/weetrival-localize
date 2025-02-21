// Package auth internal/auth/google_oauth.go
package auth

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var redirectUrl = os.Getenv("REDIRECT_URL")

var googleOauthConfig = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  fmt.Sprintf("%s%s", redirectUrl, "/auth/google/callback"),
	Scopes:       []string{"email"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin() {
	// Generate URL for Google login
	url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	log.Println("Visit the URL for Google login: ", url)
}
