package audience

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type UpdateAudienceService struct {
	authentication     auth.Authenticatable
	audienceRepository domain.AudienceRepository
	logger             grpclog.LoggerV2
}

func (s *UpdateAudienceService) Call(ctx context.Context, params *domain.UpdateAudienceParams) (*domain.UpdateAudienceResult, error) {
	if err := s.authentication.Valid(ctx); err != nil {
		return nil, err
	}

	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Errorf("[enable-audience-service] failed to retrieve a audience resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if audience == nil {
		return nil, status.Error(codes.NotFound, "Audience not found")
	}

	// Update EnabledAt by now
	// If the audience is initially disabled then enabled by this action
	if !audience.Enabled && params.Enabled {
		audience.EnabledAt = time.Now().Local()
	}

	// Update EnabledAt by zero
	// If the audience is initially enabled then disabled by this action
	if audience.Enabled && !params.Enabled {
		audience.EnabledAt = time.Time{}
	}

	audience.Enabled = params.Enabled

	if err := s.audienceRepository.Save(ctx, audience); err != nil {
		s.logger.Errorf("[enable-audience-service] failed to save a audience resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return domain.ToAudienceResult[domain.UpdateAudienceResult](audience), nil
}

func NewUpdateAudienceService(
	authentication auth.Authenticatable,
	audienceRepository domain.AudienceRepository,
	logger grpclog.LoggerV2,
) domain.AudienceUpdateable {
	return &UpdateAudienceService{
		authentication:     authentication,
		audienceRepository: audienceRepository,
		logger:             logger,
	}
}
