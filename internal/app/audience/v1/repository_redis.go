package audience

import (
	"context"
	"fmt"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	_ "github.com/google/go-querystring/query"
	"google.golang.org/grpc/grpclog"
)

const (
	REDIS_AUDIENCE_KEY             = "audiences"
	REDIS_AUDIENCE_EXPIRATION_TIME = 5 * time.Minute
)

type RedisAudienceRepository struct {
	audienceRepository domain.AudienceRepository
	db                 domain.CacheRepository
	logger             grpclog.LoggerV2
}

func (r *RedisAudienceRepository) createKey(namespace string, key interface{}) (string, error) {
	prefix := fmt.Sprintf("%s:%s", REDIS_AUDIENCE_KEY, namespace)
	return r.db.CreateKey(prefix, key)
}

func (r *RedisAudienceRepository) createPrimaryKey(name string) string {
	return fmt.Sprintf("%s/%s", REDIS_AUDIENCE_KEY, name)
}

func (r *RedisAudienceRepository) resetCache(ctx context.Context) error {
	keys := []string{}

	actionCacheKeys, err := r.db.ListKeys(ctx, REDIS_AUDIENCE_KEY+":*")
	if err != nil {
		return err
	}

	primaryCacheKeys, err := r.db.ListKeys(ctx, REDIS_AUDIENCE_KEY+"/*")
	if err != nil {
		return err
	}

	keys = append(keys, actionCacheKeys...)
	keys = append(keys, primaryCacheKeys...)

	for _, key := range keys {
		if err := r.db.Del(ctx, key); err != nil {
			return err
		}
	}

	return nil
}

// Delete implements domain.AudienceRepository
func (r *RedisAudienceRepository) Delete(ctx context.Context, audience *domain.Audience) error {
	if err := r.audienceRepository.Delete(ctx, audience); err != nil {
		return err
	}

	return r.resetCache(ctx)
}

// DeleteBy implements domain.AudienceRepository
func (r *RedisAudienceRepository) DeleteBy(ctx context.Context, args *domain.AudienceFilterArgs) error {
	if err := r.audienceRepository.DeleteBy(ctx, args); err != nil {
		return err
	}

	return r.resetCache(ctx)
}

// Get implements domain.AudienceRepository
func (r *RedisAudienceRepository) Get(ctx context.Context, featureName string, audienceId string) (audience *domain.Audience, err error) {
	key := r.createPrimaryKey(featureName)
	if err := r.db.Get(ctx, key, audience); err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to read cache: %s", err.Error())
		return nil, err
	}

	if audience != nil {
		return audience, nil
	}

	audience, err = r.audienceRepository.Get(ctx, featureName, audienceId)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to find audience: %s", err.Error())
		return nil, err
	}

	if audience == nil {
		return nil, nil
	}

	go r.db.Set(ctx, key, audience, REDIS_AUDIENCE_EXPIRATION_TIME)

	return audience, nil
}

// GetBy implements domain.AudienceRepository
func (r *RedisAudienceRepository) GetBy(ctx context.Context, args *domain.AudienceGetByArgs) (audience *domain.Audience, err error) {
	key, err := r.createKey(":get-by", args)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to create cache key: %s", err.Error())
		return nil, err
	}

	if err := r.db.Get(ctx, key, audience); err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to read cache: %s", err.Error())
		return nil, err
	}

	if audience != nil {
		return audience, nil
	}

	audience, err = r.audienceRepository.GetBy(ctx, args)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to find audience: %s", err.Error())
		return nil, err
	}

	go r.db.Set(ctx, key, audience, REDIS_AUDIENCE_EXPIRATION_TIME)

	return audience, nil
}

// List implements domain.AudienceRepository
func (r *RedisAudienceRepository) List(ctx context.Context, args *domain.AudienceListArgs) (audiences []domain.Audience, err error) {
	key, err := r.createKey("list", args)
	if err != nil {
		return nil, err
	}

	if err := r.db.Get(ctx, key, audiences); err != nil {
		return nil, err
	}

	if audiences != nil {
		return audiences, nil
	}

	audiences, err = r.audienceRepository.List(ctx, args)
	if err != nil {
		return nil, err
	}

	go r.db.Set(ctx, key, audiences, REDIS_AUDIENCE_EXPIRATION_TIME)

	return audiences, nil
}

// Save implements domain.AudienceRepository
func (r *RedisAudienceRepository) Save(ctx context.Context, audience *domain.Audience) error {
	err := r.audienceRepository.Save(ctx, audience)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to save audience on the database: %s", err.Error())
		return err
	}

	go r.db.Set(ctx, r.createPrimaryKey(audience.FeatureName), audience, REDIS_AUDIENCE_EXPIRATION_TIME)

	return nil
}

// Size implements domain.AudienceRepository
func (r *RedisAudienceRepository) Size(ctx context.Context, args *domain.AudienceFilterArgs) (uint32, error) {
	size, err := r.audienceRepository.Size(ctx, args)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to calculate audience size on the database: %s", err.Error())
		return size, err
	}

	key, err := r.createKey(":size", args)
	if err != nil {
		r.logger.Errorf("[redis-audience-repository] failed to create cache key: %s", err.Error())
		return size, err
	}

	go r.db.Set(ctx, key, size, REDIS_AUDIENCE_EXPIRATION_TIME)

	return size, nil
}

func NewRedisAudienceRepository(audienceRepository domain.AudienceRepository, db domain.CacheRepository, logger grpclog.LoggerV2) domain.AudienceRepository {
	return &RedisAudienceRepository{
		audienceRepository: audienceRepository,
		db:                 db,
		logger:             logger,
	}
}
