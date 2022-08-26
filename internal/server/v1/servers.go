package server

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/feature/v1"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

var (
	featureServer  featurev1.FeatureServiceServer
	audienceServer audiencev1.AudienceServiceServer
)

func bootstrapServers() {
	featureRepository := feature.NewPostgresRepository(db, logger)
	featureServer = feature.NewServer(
		feature.WithLogger(logger),
		feature.WithFeatureRepository(featureRepository),
	)

	audienceRepository := audience.NewPostgresRepository(db, logger)
	audienceServer = audience.NewServer(
		audience.WithAudienceRepository(audienceRepository),
		audience.WithFeatureRepository(featureRepository),
		audience.WithLogger(logger),
	)
}
