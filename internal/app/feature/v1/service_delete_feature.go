package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type DeleteFeatureService struct {
	authentication    auth.Authenticatable
	featureRepository FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *DeleteFeatureService) Call(ctx context.Context, params *DeleteParams) error {
	if err := s.authentication.Valid(ctx); err != nil {
		return err
	}

	feature, err := s.featureRepository.Get(ctx, params.Name)
	if err != nil {
		s.logger.Errorf("[delete-feature-service] failed to retrieve a feature resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return status.Error(codes.NotFound, "Feature not found")
	}

	if err := s.featureRepository.Delete(ctx, feature.Name); err != nil {
		s.logger.Errorf("[delete-feature-service] failed to delete a feature resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}

func NewDeleteFeatureService(
	authentication auth.Authenticatable,
	featureRepository FeatureRepository,
	logger grpclog.LoggerV2,
) Deletable {
	return &DeleteFeatureService{
		authentication:    authentication,
		featureRepository: featureRepository,
		logger:            logger,
	}
}
