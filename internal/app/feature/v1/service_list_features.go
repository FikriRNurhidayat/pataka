package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type ListFeaturesService struct {
	featureRepository domain.FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *ListFeaturesService) Call(ctx context.Context, params *domain.ListFeaturesParams) (*domain.ListFeaturesResult, error) {
	filter := &domain.FeatureFilterArgs{
		Q:       params.Q,
		Enabled: params.Enabled,
	}

	features, err := s.featureRepository.List(ctx, &domain.FeatureListArgs{
		Limit:  params.GetLimit(),
		Offset: params.GetOffset(),
		Sort:   params.Sort,
		Filter: filter,
	})

	if err != nil {
		s.logger.Error("[list-features-service] failed to list feature collection")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	featureSize, err := s.featureRepository.Size(ctx, filter)
	if err != nil {
		s.logger.Error("[list-features-service] failed to measure feature collection size")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &domain.ListFeaturesResult{
		PaginationResult: params.PaginationResult(featureSize),
		Size:             featureSize,
		Features:         features,
	}, nil
}

func NewListFeaturesService(featureRepository domain.FeatureRepository, logger grpclog.LoggerV2) domain.FeaturesListable {
	return &ListFeaturesService{
		featureRepository: featureRepository,
		logger:            logger,
	}
}
