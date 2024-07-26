package config

import (
	"os"
	"time"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_CONN"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		IdleTimeout: 5 * time.Minute,
		MaxRetries:  2,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}