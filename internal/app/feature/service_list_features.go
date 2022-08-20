package feature

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type Listable interface {
	Call(context.Context, *ListParams) (*ListResult, error)
}

type ListFeaturesService struct {
	FeatureRepository FeatureRepository
	Logger            grpclog.LoggerV2
	DefaultPageNumber uint32
	DefaultPageSize   uint32
}

func (s *ListFeaturesService) Call(ctx context.Context, params *ListParams) (*ListResult, error) {
	filter := &FeatureFilterArgs{
		Q:       params.Q,
		Enabled: params.Enabled,
	}

	features, err := s.FeatureRepository.List(ctx, &FeatureListArgs{
		Limit:  params.ToLimit(s.DefaultPageSize),
		Offset: params.ToOffset(s.DefaultPageNumber),
		Sort:   params.Sort,
		Filter: filter,
	})

	if err != nil {
		s.Logger.Error("[list-features-service] failed to list feature collection")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	featureSize, err := s.FeatureRepository.Size(ctx, filter)
	if err != nil {
		s.Logger.Error("[list-features-service] failed to measure feature collection size")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &ListResult{
		PaginationResult: params.ToPaginationResult(featureSize),
		Size:             featureSize,
		Features:         features,
	}, nil
}

func NewListFeaturesService(
	FeatureRepository FeatureRepository,
	Logger grpclog.LoggerV2,
	DefaultPageNumber uint32,
	DefaultPageSize uint32,
) Listable {
	return &ListFeaturesService{
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
		DefaultPageNumber: DefaultPageNumber,
		DefaultPageSize:   DefaultPageSize,
	}
}
