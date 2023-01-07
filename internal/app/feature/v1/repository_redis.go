package feature

import (
	"context"
	"fmt"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	_ "github.com/google/go-querystring/query"
	"google.golang.org/grpc/grpclog"
)

const (
	REDIS_FEATURE_KEY             = "features"
	REDIS_FEATURE_EXPIRATION_TIME = 5 * time.Minute
)

type RedisFeatureRepository struct {
	featureRepository domain.FeatureRepository
	db                domain.CacheRepository
	logger            grpclog.LoggerV2
}

func (r *RedisFeatureRepository) createKey(namespace string, key interface{}) (string, error) {
	prefix := fmt.Sprintf("%s:%s", REDIS_FEATURE_KEY, namespace)
	return r.db.CreateKey(prefix, key)
}

func (r *RedisFeatureRepository) createPrimaryKey(name string) string {
	return fmt.Sprintf("%s/%s", REDIS_FEATURE_KEY, name)
}

func (r *RedisFeatureRepository) resetCache(ctx context.Context) error {
	keys := []string{}

	actionCacheKeys, err := r.db.ListKeys(ctx, REDIS_FEATURE_KEY+":*")
	if err != nil {
		return err
	}

	primaryCacheKeys, err := r.db.ListKeys(ctx, REDIS_FEATURE_KEY+"/*")
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

// Delete implements domain.FeatureRepository
func (r *RedisFeatureRepository) Delete(ctx context.Context, name string) error {
	if err := r.featureRepository.Delete(ctx, name); err != nil {
		return err
	}

	return r.resetCache(ctx)
}

// DeleteBy implements domain.FeatureRepository
func (r *RedisFeatureRepository) DeleteBy(ctx context.Context, args *domain.FeatureFilterArgs) error {
	if err := r.featureRepository.DeleteBy(ctx, args); err != nil {
		return err
	}

	return r.resetCache(ctx)
}

// Get implements domain.FeatureRepository
func (r *RedisFeatureRepository) Get(ctx context.Context, name string) (feature *domain.Feature, err error) {
	key := r.createPrimaryKey(name)
	if err := r.db.Get(ctx, key, feature); err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to read cache: %s", err.Error())
		return nil, err
	}

	if feature != nil {
		return feature, nil
	}

	feature, err = r.featureRepository.Get(ctx, name)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to find feature: %s", err.Error())
		return nil, err
	}

	if feature == nil {
		return nil, nil
	}

	go r.db.Set(ctx, key, feature, REDIS_FEATURE_EXPIRATION_TIME)

	return feature, nil
}

// GetBy implements domain.FeatureRepository
func (r *RedisFeatureRepository) GetBy(ctx context.Context, args *domain.FeatureGetByArgs) (feature *domain.Feature, err error) {
	key, err := r.createKey(":get-by", args)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to create cache key: %s", err.Error())
		return nil, err
	}

	if err := r.db.Get(ctx, key, feature); err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to read cache: %s", err.Error())
		return nil, err
	}

	if feature != nil {
		return feature, nil
	}

	feature, err = r.featureRepository.GetBy(ctx, args)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to find feature: %s", err.Error())
		return nil, err
	}

	go r.db.Set(ctx, key, feature, REDIS_FEATURE_EXPIRATION_TIME)

	return feature, nil
}

// List implements domain.FeatureRepository
func (r *RedisFeatureRepository) List(ctx context.Context, args *domain.FeatureListArgs) (features []domain.Feature, err error) {
	key, err := r.createKey("list", args)
	if err != nil {
		return nil, err
	}

	if err := r.db.Get(ctx, key, features); err != nil {
		return nil, err
	}

	if features != nil {
		return features, nil
	}

	features, err = r.featureRepository.List(ctx, args)
	if err != nil {
		return nil, err
	}

	go r.db.Set(ctx, key, features, REDIS_FEATURE_EXPIRATION_TIME)

	return features, nil
}

// Save implements domain.FeatureRepository
func (r *RedisFeatureRepository) Save(ctx context.Context, feature *domain.Feature) error {
	err := r.featureRepository.Save(ctx, feature)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to save feature on the database: %s", err.Error())
		return err
	}

	go r.db.Set(ctx, r.createPrimaryKey(feature.Name), feature, REDIS_FEATURE_EXPIRATION_TIME)

	return nil
}

// Size implements domain.FeatureRepository
func (r *RedisFeatureRepository) Size(ctx context.Context, args *domain.FeatureFilterArgs) (uint32, error) {
	size, err := r.featureRepository.Size(ctx, args)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to calculate feature size on the database: %s", err.Error())
		return size, err
	}

	key, err := r.createKey(":size", args)
	if err != nil {
		r.logger.Errorf("[redis-feature-repository] failed to create cache key: %s", err.Error())
		return size, err
	}

	go r.db.Set(ctx, key, size, REDIS_FEATURE_EXPIRATION_TIME)

	return size, nil
}

func NewRedisFeatureRepository(featureRepository domain.FeatureRepository, db domain.CacheRepository, logger grpclog.LoggerV2) domain.FeatureRepository {
	return &RedisFeatureRepository{
		featureRepository: featureRepository,
		db:                db,
		logger:            logger,
	}
}
