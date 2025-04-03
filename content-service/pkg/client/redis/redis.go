package redis

import (
	"github.com/redis/go-redis/v9"
	"net"
)

func NewClient(host, port string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(host, port),
		Password: "",
		DB:       0,
	})

	return rdb
}
