package dao

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	return rc.rdb.Set(ctx, key, value, expire).Err()
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.203.29:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	Rc = &RedisCache{rdb: rdb}
}
