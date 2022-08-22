package feature

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type ListFeaturesService struct {
	featureRepository FeatureRepository
	logger            grpclog.LoggerV2
	defaultPageNumber uint32
	defaultPageSize   uint32
}

func (s *ListFeaturesService) Call(ctx context.Context, params *ListParams) (*ListResult, error) {
	filter := &FeatureFilterArgs{
		Q:       params.Q,
		Enabled: params.Enabled,
	}

	features, err := s.featureRepository.List(ctx, &FeatureListArgs{
		Limit:  params.ToLimit(s.defaultPageSize),
		Offset: params.ToOffset(s.defaultPageNumber),
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

	return &ListResult{
		PaginationResult: params.ToPaginationResult(featureSize),
		Size:             featureSize,
		Features:         features,
	}, nil
}

func NewListFeaturesService(
	featureRepository FeatureRepository,
	logger grpclog.LoggerV2,
	defaultPageNumber uint32,
	defaultPageSize uint32,
) Listable {
	return &ListFeaturesService{
		featureRepository: featureRepository,
		logger:            logger,
		defaultPageNumber: defaultPageNumber,
		defaultPageSize:   defaultPageSize,
	}
}