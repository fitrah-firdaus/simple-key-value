package configuration

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Remove(key string) error
}

type redisCache struct {
	rdb *redis.Client
}

func (r redisCache) Set(key string, value string) error {
	return r.rdb.Set(context.Background(), key, value, 60*time.Second).Err()
}

func (r redisCache) Get(key string) (string, error) {
	return r.rdb.Get(context.Background(), key).Result()
}

func (r redisCache) Remove(key string) error {
	return r.rdb.Del(context.Background(), key).Err()
}

func NewRedisCache(config Config) RedisCache {
	url := config.Get("REDIS_URI")
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Error(err)
	}
	return &redisCache{rdb: redis.NewClient(opts)}
}
