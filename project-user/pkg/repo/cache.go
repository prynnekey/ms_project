package repo

import (
	"time"

	"golang.org/x/net/context"
)

// 缓存接口
type Cache interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
