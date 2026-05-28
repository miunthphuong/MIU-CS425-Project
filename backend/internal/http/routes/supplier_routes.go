package routes

import (
	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/http/middleware"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func registerSupplierRoutes(api *gin.RouterGroup, jwtService auth.JWTService) {
	supplier := api.Group("/supplier")
	supplier.Use(middleware.RequireAuth(jwtService))
	supplier.Use(middleware.RequireRole(models.RoleSupplier))
	supplier.GET("/me", currentUser)
	supplier.GET("/dashboard", roleDashboard(models.RoleSupplier))
}
