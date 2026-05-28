package routes

import (
	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/http/middleware"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func registerCustomerRoutes(api *gin.RouterGroup, jwtService auth.JWTService) {
	customer := api.Group("/customer")
	customer.Use(middleware.RequireAuth(jwtService))
	customer.Use(middleware.RequireRole(models.RoleCustomer))
	customer.GET("/me", currentUser)
	customer.GET("/dashboard", roleDashboard(models.RoleCustomer))
}
