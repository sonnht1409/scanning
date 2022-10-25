package cache

import (
	"context"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v9"
	"github.com/sonnht1409/scanning/service/config"
)

func NewCache() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Values.Redis.URL,
		DB:   config.Values.Redis.DB,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return redisClient
}

func NewCacheLock(client *redis.Client) *redislock.Client {
	locker := redislock.New(client)
	return locker
}
