package database

import (
	"finalAssing/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		// Addr: "redis:6379",
		// Password: "",
		// DB: 0,
		Addr:     cfg.RedisConfig.RedisAddr,
		Password: cfg.RedisConfig.RedisPassword,
		DB:       cfg.RedisConfig.RedisDb,
	})
	return client
}
