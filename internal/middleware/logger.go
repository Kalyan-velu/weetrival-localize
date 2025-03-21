package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		log.Printf("Request: %s %s\nHeaders: %v\nBody: %s",
			c.Request.Method,
			c.Request.URL,
			c.Request.Header,
			string(body),
		)

		// Reset request body for the next handler
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}
}
