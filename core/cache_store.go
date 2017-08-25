package core

import (
	"time"

	"github.com/go-redis/redis"
)

// CacheStore simple redis implementation
type CacheStore struct {
	Cache *redis.Client
}

func (s *CacheStore) Ping() error {
	return s.Cache.Ping().Err()
}

func (s *CacheStore) Get(key string) (string, error) {
	return s.Cache.Get(key).Result()
}

func (s *CacheStore) Set(key string, value interface{}, exp time.Duration) (string, error) {
	return s.Cache.Set(key, value, exp).Result()
}
