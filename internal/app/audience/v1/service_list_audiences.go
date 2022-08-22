package audience

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type ListAudiencesService struct {
	audienceRepository AudienceRepository
	logger             grpclog.LoggerV2
	defaultPageNumber  uint32
	defaultPageSize    uint32
}

func (s *ListAudiencesService) Call(ctx context.Context, params *ListParams) (*ListResult, error) {
	filter := &AudienceFilterArgs{
		FeatureName: params.FeatureName,
		AudienceIds: params.AudienceIds,
		Enabled:     params.Enabled,
	}

	audiences, err := s.audienceRepository.List(ctx, &AudienceListArgs{
		Limit:  params.ToLimit(s.defaultPageSize),
		Offset: params.ToOffset(s.defaultPageNumber),
		Sort:   params.Sort,
		Filter: filter,
	})

	if err != nil {
		s.logger.Error("[list-audiences-service] failed to list audience collection")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	audienceSize, err := s.audienceRepository.Size(ctx, filter)
	if err != nil {
		s.logger.Error("[list-audiences-service] failed to measure audience collection size")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &ListResult{
		PaginationResult: params.ToPaginationResult(audienceSize),
		Size:             audienceSize,
		Audiences:        audiences,
	}, nil
}

func NewListAudiencesService(
	audienceRepository AudienceRepository,
	logger grpclog.LoggerV2,
	defaultPageNumber uint32,
	defaultPageSize uint32,
) Listable {
	return &ListAudiencesService{
		audienceRepository: audienceRepository,
		logger:             logger,
		defaultPageNumber:  defaultPageNumber,
		defaultPageSize:    defaultPageSize,
	}
}
