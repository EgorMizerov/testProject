package repository

import (
	"fmt"
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
	fmt.Println(addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       0,
	})

	return rdb
}
