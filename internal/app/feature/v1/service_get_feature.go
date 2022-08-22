package feature

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type GetFeatureService struct {
	featureRepository FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *GetFeatureService) Call(ctx context.Context, params *GetParams) (*GetResult, error) {
	feature, err := s.featureRepository.Get(ctx, params.Name)
	if err != nil {
		s.logger.Error("[get-feature-service] failed to retrieve a feature resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return nil, status.Error(codes.NotFound, "Feature not found")
	}

	return ToFeatureResult[GetResult](feature), nil
}

func NewGetFeatureService(
	featureRepository FeatureRepository,
	logger grpclog.LoggerV2,
) Getable {
	return &GetFeatureService{
		featureRepository: featureRepository,
		logger:            logger,
	}
}
