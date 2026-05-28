package routes

import (
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func registerRoleAuthRoutes(api *gin.RouterGroup, authHandler AuthHandler) {
	adminAuth := api.Group("/admin/auth")
	adminAuth.POST("/register", authHandler.RegisterForRole(models.RoleAdmin))
	adminAuth.POST("/login", authHandler.LoginForRole(models.RoleAdmin))

	supplierAuth := api.Group("/supplier/auth")
	supplierAuth.POST("/register", authHandler.RegisterForRole(models.RoleSupplier))
	supplierAuth.POST("/login", authHandler.LoginForRole(models.RoleSupplier))

	customerAuth := api.Group("/customer/auth")
	customerAuth.POST("/register", authHandler.RegisterForRole(models.RoleCustomer))
	customerAuth.POST("/login", authHandler.LoginForRole(models.RoleCustomer))
}
