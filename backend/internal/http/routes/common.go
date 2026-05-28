package routes

import (
	"net/http"

	"miu-cs425-project/backend/internal/http/middleware"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func currentUser(c *gin.Context) {
	role, _ := c.Get(middleware.CurrentRoleKey)

	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetUint(middleware.CurrentUserIDKey),
		"role":    role,
	})
}

func roleDashboard(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": string(role) + " dashboard",
		})
	}
}
