package audience

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type UpdateAudienceService struct {
	unitOfWork domain.UnitOfWork
	logger     grpclog.LoggerV2
}

func (s *UpdateAudienceService) Call(ctx context.Context, params *domain.UpdateAudienceParams) (*domain.UpdateAudienceResult, error) {
	var audience *domain.Audience

	err := s.unitOfWork.Do(ctx, func(r domain.Repository) (err error) {
		audience, err = r.AudienceRepository().Get(ctx, params.FeatureName, params.AudienceId)
		if err != nil {
			s.logger.Errorf("[update-audience-service] failed to retrieve a audience resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		if audience == nil {
			return status.Error(codes.NotFound, "Audience not found")
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

		if err = r.AudienceRepository().Save(ctx, audience); err != nil {
			s.logger.Errorf("[update-audience-service] failed to save a audience resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return domain.ToAudienceResult[domain.UpdateAudienceResult](audience), nil
}

func NewUpdateAudienceService(
	unitOfWork domain.UnitOfWork,
	logger grpclog.LoggerV2,
) domain.AudienceUpdateable {
	return &UpdateAudienceService{
		unitOfWork: unitOfWork,
		logger:     logger,
	}
}
