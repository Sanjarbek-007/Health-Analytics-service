package rediss

import (
	"github.com/redis/go-redis/v9"
)

func ConnectRDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "root",
		DB:       0,
	})

	return rdb
}
