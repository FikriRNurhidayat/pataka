package audience

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type GetAudienceService struct {
	audienceRepository AudienceRepository
	logger             grpclog.LoggerV2
}

func (s *GetAudienceService) Call(ctx context.Context, params *GetParams) (*GetResult, error) {
	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Error("[get-audience-service] failed to retrieve a audience resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if audience == nil {
		return nil, status.Error(codes.NotFound, "Audience not found")
	}

	return ToAudienceResult[GetResult](audience), nil
}

func NewGetAudienceService(
	audienceRepository AudienceRepository,
	logger grpclog.LoggerV2,
) Getable {
	return &GetAudienceService{
		audienceRepository: audienceRepository,
		logger:             logger,
	}
}
