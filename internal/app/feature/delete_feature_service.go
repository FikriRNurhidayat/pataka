package feature

import (
	"context"
	"errors"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/grpclog"
)

type Deletable interface {
	Call(context.Context, *DeleteParams) error
}

type DeleteFeatureService struct {
	FeatureRepository domain.FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *DeleteFeatureService) Call(ctx context.Context, params *DeleteParams) error {
	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to retrieve a feature resource: %s", err.Error())
		return err
	}

	if feature == nil {
		return errors.New("Feature does not exist")
	}

	if err := s.FeatureRepository.Delete(ctx, feature.Name); err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to delete a feature resource: %s", err.Error())
		return err
	}

	return nil
}

func NewDeleteFeatureService(
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Deletable {
	return &DeleteFeatureService{
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
