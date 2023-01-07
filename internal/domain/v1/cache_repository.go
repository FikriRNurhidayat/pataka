package domain

import (
	"context"
	"time"
)

type CacheRepository interface {
	CreateKey(prefix string, key interface{}) (string, error)
	ListKeys(ctx context.Context, match string) ([]string, error)
	Get(ctx context.Context, k string, o any) error
	Set(ctx context.Context, k string, v interface{}, e time.Duration) (err error)
	Del(ctx context.Context, k string) error
}
