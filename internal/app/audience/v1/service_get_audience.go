package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type GetAudienceService struct {
	audienceRepository domain.AudienceRepository
	logger             grpclog.LoggerV2
}

func (s *GetAudienceService) Call(ctx context.Context, params *domain.GetAudienceParams) (*domain.GetAudienceResult, error) {
	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Error("[get-audience-service] failed to retrieve a audience resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if audience == nil {
		return nil, status.Error(codes.NotFound, "Audience not found")
	}

	return domain.ToAudienceResult[domain.GetAudienceResult](audience), nil
}

func NewGetAudienceService(
	audienceRepository domain.AudienceRepository,
	logger grpclog.LoggerV2,
) domain.AudienceGetable {
	return &GetAudienceService{
		audienceRepository: audienceRepository,
		logger:             logger,
	}
}
