package core

import (
	"github.com/go-redis/redis"
)

type CacheStore struct {
	Cache *redis.Client
}

func (s *CacheStore) Ping() error {
	return s.Cache.Ping().Err()
}
