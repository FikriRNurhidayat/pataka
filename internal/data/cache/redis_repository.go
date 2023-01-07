package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/go-redis/redis/v8"
	"github.com/google/go-querystring/query"
	"google.golang.org/grpc/grpclog"
)

type redisRepository struct {
	client redis.UniversalClient
	logger grpclog.LoggerV2
}

// ListKeys implements domain.CacheRepository
func (r *redisRepository) ListKeys(ctx context.Context, prefix string) ([]string, error) {
	var cursor uint64
	var ks []string
	for {
		keys, cursor, err := r.client.Scan(ctx, cursor, prefix, 0).Result()
		if err != nil {
			return nil, err
		}

		for _, k := range keys {
			ks = append(ks, k)
		}

		if cursor == 0 {
			break
		}
	}

	return ks, nil
}

// CreateKey implements domain.CacheRepository
func (*redisRepository) CreateKey(prefix string, key interface{}) (string, error) {
	value, err := query.Values(key)
	if err != nil {
		return "", err
	}

	return prefix + value.Encode(), nil
}

// Del implements RedisRepository
func (r *redisRepository) Del(ctx context.Context, k string) error {
	cmd := r.client.Del(ctx, k)
	if err := cmd.Err(); err != nil {
		return err
	}

	r.logger.Infof("[redis-repository] delete cache: %s", k)
	return nil
}

func (r *redisRepository) Get(ctx context.Context, k string, o any) error {
	value, err := r.client.Get(ctx, k).Result()

	if err == redis.Nil {
		r.logger.Infof("[redis-repository] cache not found: %s", k)
		*&o = nil
		return nil
	}

	if err != nil {
		*&o = nil
		return err
	}

	r.logger.Infof("[redis-repository] parsing cache: %s", k)
	return json.Unmarshal([]byte(value), o)
}

func (r *redisRepository) Set(ctx context.Context, k string, v interface{}, e time.Duration) (err error) {
	value, err := json.Marshal(v)

	if err != nil {
		r.logger.Warning("[redis-repository] fail to write cache: %s", k)
		return err
	}

	r.logger.Infof("[redis-repository] writing cache: %s", k)
	return r.client.Set(ctx, k, string(value), e).Err()
}

func NewRedisRepository(client redis.UniversalClient, logger grpclog.LoggerV2) domain.CacheRepository {
	return &redisRepository{
		client: client,
		logger: logger,
	}
}
