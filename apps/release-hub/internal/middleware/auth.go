package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maodou24/release-hub/internal/pkg/security"
)

const userIDContextKey = "authenticatedUserID"

// SessionAuth optionally validates the signed session cookie and stores its
// user ID in the Gin context. Missing or invalid sessions do not block requests;
// handlers that require authentication decide how to respond.
func SessionAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(security.UserSessionCookie)
		if err == nil {
			if userID, parseErr := security.ParseUserSession(cookie.Value, secret); parseErr == nil {
				c.Set(userIDContextKey, userID)
			}
		}
		c.Next()
	}
}

// UserID returns the authenticated user ID placed in the context by SessionAuth.
func UserID(c *gin.Context) (string, bool) {
	value, ok := c.Get(userIDContextKey)
	if !ok {
		return "", false
	}
	userID, ok := value.(string)
	return userID, ok && userID != ""
}
