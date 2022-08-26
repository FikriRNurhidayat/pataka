package transaction

import (
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
	"google.golang.org/grpc/grpclog"
)

type repository struct {
	featureRepository  domain.FeatureRepository
	audienceRepository domain.AudienceRepository
}

func (u *repository) AudienceRepository() domain.AudienceRepository {
	return u.audienceRepository
}

func (u *repository) FeatureRepository() domain.FeatureRepository {
	return u.featureRepository
}

func makeRepository(db driver.DB, logger grpclog.LoggerV2, factory *RepositoryFactory) domain.Repository {
	return &repository{
		featureRepository:  factory.FeatureRepository(db, logger),
		audienceRepository: factory.AudienceRepository(db, logger),
	}
}
