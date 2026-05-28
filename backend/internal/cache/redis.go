package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Connect(ctx context.Context, addr, password string, db int, log *zap.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	log.Info("connected to Redis", zap.String("addr", addr), zap.Int("db", db))
	return client, nil
}
