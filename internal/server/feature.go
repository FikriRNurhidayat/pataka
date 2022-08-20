package server

import (
	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/feature"
	"github.com/fikrirnurhidayat/ffgo/internal/domain"

	"github.com/fikrirnurhidayat/ffgo/internal/app/feature/repository"
)

var (
	featureCreator    feature.Createable
	featureLister     feature.Listable
	featureGetter     feature.Getable
	featureUpdater    feature.Updatable
	featureDeletor    feature.Deletable
	featureRepository domain.FeatureRepository
	featureServer     featureflag.FeatureServiceServer
)

func initFeatureServer() {
	featureRepository = feature_repository.NewPostgresFeatureRepository(db, logger)
	featureCreator = feature.NewCreateFeatureService(authenticationService, featureRepository, logger)
	featureLister = feature.NewListFeaturesService(featureRepository, logger, 1, 10)
	featureGetter = feature.NewGetFeatureService(featureRepository, logger)
	featureUpdater = feature.NewUpdateFeatureService(authenticationService, featureRepository, logger)
	featureDeletor = feature.NewDeleteFeatureService(authenticationService, featureRepository, logger)

	featureServer = feature.NewFeatureServer(
		feature.WithCreator(featureCreator),
		feature.WithUpdater(featureUpdater),
		feature.WithDeletor(featureDeletor),
		feature.WithLister(featureLister),
		feature.WithGetter(featureGetter),
		feature.WithLogger(logger),
	)
}
