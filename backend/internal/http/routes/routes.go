package routes

import (
	"miu-cs425-project/backend/internal/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, jwtService auth.JWTService) {
	authHandler := NewAuthHandler(db, jwtService)

	api := router.Group("/api/v1")
	api.GET("/health", health)

	registerRoleAuthRoutes(api, authHandler)

	registerAdminRoutes(api, db, jwtService)
	registerSupplierRoutes(api, jwtService)
	registerCustomerRoutes(api, jwtService)
}
