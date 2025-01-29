package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_URL"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})

	return rdb
}
