package server

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/feature/v1"
	uow "github.com/fikrirnurhidayat/ffgo/internal/data/atomic"
	"github.com/fikrirnurhidayat/ffgo/internal/data/cache"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

var (
	featureServer  featurev1.FeatureServiceServer
	audienceServer audiencev1.AudienceServiceServer
)

func bootstrapServers() {
	redisRepository := cache.NewRedisRepository(redisClient, logger)

	featurePostgresRepository := feature.NewPostgresRepository(db, logger)
	audiencePostgresRepository := audience.NewPostgresRepository(db, logger)

	featureRedisRepository := feature.NewRedisFeatureRepository(featurePostgresRepository, redisRepository, logger)
	audienceRedisRepository := audience.NewRedisAudienceRepository(audiencePostgresRepository, redisRepository, logger)

	// TODO: Also adds Redis on the Unit Of Work
	// 		 Maybe, it's viable, maybe it's not.
	tx := uow.NewPostgresUnitOfWork(db, logger, &uow.PostgresRepositoryFactory{
		FeatureRepository:  feature.NewPostgresRepository,
		AudienceRepository: audience.NewPostgresRepository,
	})

	featureServer = feature.NewServer(
		feature.WithLogger(logger),
		feature.WithUnitOfWork(tx),
		feature.WithFeatureRepository(featureRedisRepository),
	)

	audienceServer = audience.NewServer(
		audience.WithAudienceRepository(audienceRedisRepository),
		audience.WithUnitOfWork(tx),
		audience.WithLogger(logger),
	)
}
