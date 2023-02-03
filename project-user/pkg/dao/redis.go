package dao

import (
	"context"
	"time"

	"github.com/prynnekey/ms_project/project-user/config"
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
	rdb := redis.NewClient(config.AppConfig.ReadRedisConfig())
	Rc = &RedisCache{rdb: rdb}
}
