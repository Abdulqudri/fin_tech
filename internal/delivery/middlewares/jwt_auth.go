package middleware

import (
	"net/http"
	"strings"

	"github.com/Abdulqudri/fintech/internal/utils/security"
	"github.com/gin-gonic/gin"
)

func JWTAuth(jwt *security.JWTIssuer) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		userID, err := jwt.VerifyAccessToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Inject into context for handlers/services
		c.Set("user_id", userID)

		c.Next()
	}
}
