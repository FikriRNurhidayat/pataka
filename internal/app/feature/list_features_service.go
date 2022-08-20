package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/grpclog"
)

type Listable interface {
	Call(context.Context, *ListParams) (*ListResult, error)
}

type ListFeaturesService struct {
	FeatureRepository domain.FeatureRepository
	Logger            grpclog.LoggerV2
	DefaultPageNumber uint32
	DefaultPageSize   uint32
}

func (s *ListFeaturesService) Call(ctx context.Context, params *ListParams) (*ListResult, error) {
	filter := &domain.FeatureFilterArgs{
		Q:       params.Q,
		Enabled: params.Enabled,
	}

	features, err := s.FeatureRepository.List(ctx, &domain.FeatureListArgs{
		Limit:  params.ToLimit(s.DefaultPageSize),
		Offset: params.ToOffset(s.DefaultPageNumber),
		Sort:   params.Sort,
		Filter: filter,
	})

	if err != nil {
		s.Logger.Infof("[FeatureRepository] failed to list feature collection: %v", err.Error())
		return nil, err
	}

	featureSize, err := s.FeatureRepository.Size(ctx, filter)
	if err != nil {
		s.Logger.Infof("[FeatureRepository] failed to measure feature collection size: %v", err.Error())
		return nil, err
	}

	return &ListResult{
		PaginationResult: params.ToPaginationResult(featureSize),
		Size:             featureSize,
		Features:         features,
	}, nil
}

func NewListFeaturesService(
	FeatureRepository domain.FeatureRepository,
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
