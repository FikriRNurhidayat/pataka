package server

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/feature/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/transaction"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

var (
	featureServer  featurev1.FeatureServiceServer
	audienceServer audiencev1.AudienceServiceServer
)

func bootstrapServers() {
	featureRepository := feature.NewPostgresRepository(db, logger)
	audienceRepository := audience.NewPostgresRepository(db, logger)

	tx := transaction.New(db, logger, &transaction.RepositoryFactory{
		FeatureRepository:  feature.NewPostgresRepository,
		AudienceRepository: audience.NewPostgresRepository,
	})

	featureServer = feature.NewServer(
		feature.WithLogger(logger),
		feature.WithUnitOfWork(tx),
		feature.WithFeatureRepository(featureRepository),
	)

	audienceServer = audience.NewServer(
		audience.WithAudienceRepository(audienceRepository),
		audience.WithUnitOfWork(tx),
		audience.WithLogger(logger),
	)
}
