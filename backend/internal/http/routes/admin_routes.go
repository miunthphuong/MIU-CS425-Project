package routes

import (
	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/http/middleware"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerAdminRoutes(api *gin.RouterGroup, db *gorm.DB, jwtService auth.JWTService) {
	userHandler := NewUserHandler(db)

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAuth(jwtService))
	admin.Use(middleware.RequireRole(models.RoleAdmin))
	admin.GET("/me", currentUser)
	admin.GET("/dashboard", roleDashboard(models.RoleAdmin))
	admin.GET("/users", userHandler.List)
	admin.POST("/users", userHandler.Create)
	admin.GET("/users/:id", userHandler.Get)
	admin.PUT("/users/:id", userHandler.Update)
	admin.DELETE("/users/:id", userHandler.Delete)
}
