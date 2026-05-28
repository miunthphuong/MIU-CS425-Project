package middleware

import (
	"net/http"
	"strings"

	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	CurrentUserIDKey = "current_user_id"
	CurrentRoleKey   = "current_role"
)

func RequireAuth(jwtService auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := bearerToken(c.GetHeader("Authorization"))
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}

		claims, err := jwtService.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set(CurrentUserIDKey, claims.UserID)
		c.Set(CurrentRoleKey, claims.Role)
		c.Next()
	}
}

func RequireRole(roles ...models.Role) gin.HandlerFunc {
	allowedRoles := make(map[models.Role]struct{}, len(roles))
	for _, role := range roles {
		allowedRoles[role] = struct{}{}
	}

	return func(c *gin.Context) {
		role, ok := c.Get(CurrentRoleKey)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authenticated role"})
			return
		}

		userRole, ok := role.(models.Role)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid authenticated role"})
			return
		}

		if _, exists := allowedRoles[userRole]; !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}

func bearerToken(header string) string {
	const prefix = "Bearer "

	if !strings.HasPrefix(header, prefix) {
		return ""
	}

	return strings.TrimSpace(strings.TrimPrefix(header, prefix))
}
