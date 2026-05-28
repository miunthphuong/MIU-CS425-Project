package main

import (
	"context"
	"log"

	"miu-cs425-project/backend/internal/cache"
	"miu-cs425-project/backend/internal/config"
	"miu-cs425-project/backend/internal/database"
	"miu-cs425-project/backend/internal/server"
	"miu-cs425-project/backend/internal/storage"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables")
	}

	cfg := config.Load()

	appLogger, err := zap.NewProduction()
	if cfg.AppEnv == "development" {
		appLogger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("could not initialize logger: %v", err)
	}
	defer appLogger.Sync()

	db, err := database.Connect(cfg.DatabaseURL, appLogger)
	if err != nil {
		appLogger.Fatal("could not connect to database", zap.Error(err))
	}

	ctx := context.Background()
	redisClient, err := cache.Connect(ctx, cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB, appLogger)
	if err != nil {
		appLogger.Fatal("could not connect to Redis", zap.Error(err))
	}
	defer redisClient.Close()

	_, err = storage.Connect(ctx, cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOBucket, cfg.MinIOUseSSL, appLogger)
	if err != nil {
		appLogger.Fatal("could not connect to MinIO", zap.Error(err))
	}

	router := server.New(cfg, db, appLogger)
	if err := router.Run(":" + cfg.Port); err != nil {
		appLogger.Fatal("server stopped", zap.Error(err))
	}
}
