// internal/middleware/abac.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ABAC middleware checks user's role and permission
func ABAC(c *gin.Context) {
	userRole := c.GetHeader("Role") // Simulating getting user role from the request (could be from JWT token)
	if userRole != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}
	c.Next() // Allow request to continue if authorized
}
