package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
)

type Config struct {
	Host     string
	Port     string
	Password string
}

func ConnectRedis(cfg Config) *redis.Client {
	addr := strings.Join([]string{cfg.Host, cfg.Port}, ":")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       0,
	})
	err := rdb.Ping(context.TODO()).Err()
	if err != nil {
		panic(err.Error())
	}

	return rdb
}
