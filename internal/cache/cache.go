package cache

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/starptech/go-web/config"
)

func NewCache(config *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPwd,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil || pong == "" {
		log.Fatalf("redis cache: got no PONG back %q", err)
	}

	return client
}
