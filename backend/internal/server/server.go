package server

import (
	"time"

	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/config"
	"miu-cs425-project/backend/internal/http/middleware"
	"miu-cs425-project/backend/internal/http/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func New(cfg config.Config, db *gorm.DB, log *zap.Logger) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger(log))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.TokenDuration)
	routes.Register(router, db, jwtService)

	return router
}
