package server

import (
	"log"

	"github.com/go-redis/redis"
)

func NewCache(addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil || pong == "" {
		log.Fatalf("redis cache: got no PONG back %q", err)
	}

	return client
}
