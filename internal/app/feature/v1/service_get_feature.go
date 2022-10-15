package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type GetFeatureService struct {
	featureRepository domain.FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *GetFeatureService) Call(ctx context.Context, params *domain.GetFeatureParams) (*domain.GetFeatureResult, error) {
	feature, err := s.featureRepository.Get(ctx, params.Name)
	if err != nil {
		s.logger.Error("[get-feature-service] failed to retrieve a feature resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return nil, status.Error(codes.NotFound, "Feature not found")
	}

	return domain.ToFeatureResult[domain.GetFeatureResult](feature), nil
}

func NewGetFeatureService(featureRepository domain.FeatureRepository, logger grpclog.LoggerV2) domain.FeatureGetable {
	return &GetFeatureService{
		featureRepository: featureRepository,
		logger:            logger,
	}
}
