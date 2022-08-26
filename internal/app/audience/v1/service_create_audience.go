package audience

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type CreateAudienceService struct {
	unitOfWork domain.UnitOfWork
	logger     grpclog.LoggerV2
}

func (s *CreateAudienceService) Call(ctx context.Context, params *domain.CreateAudienceParams) (*domain.CreateAudienceResult, error) {
	var audience *domain.Audience

	err := s.unitOfWork.Do(ctx, func(r domain.Repository) error {
		feature, err := r.FeatureRepository().Get(ctx, params.FeatureName)
		if err != nil {
			s.logger.Error("[create-audience-service] failed to retrieve feature on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		if feature == nil {
			return status.Error(codes.NotFound, "Feature not found")
		}

		audience, err = r.AudienceRepository().Get(ctx, params.FeatureName, params.AudienceId)
		if err != nil {
			s.logger.Error("[create-audience-service] failed to retrieve audience on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		if audience != nil {
			return status.Error(codes.AlreadyExists, "Audience already exist")
		}

		audience = &domain.Audience{
			FeatureName: params.FeatureName,
			AudienceId:  params.AudienceId,
			Enabled:     params.Enabled,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if params.Enabled {
			audience.EnabledAt = time.Now().Local()
		}

		if err = r.AudienceRepository().Save(ctx, audience); err != nil {
			s.logger.Error("[create-audience-service] failed to save audience on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		feature.HasAudience = true

		if err = r.FeatureRepository().Save(ctx, feature); err != nil {
			s.logger.Error("[create-audience-service] failed to save feature on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return domain.ToAudienceResult[domain.CreateAudienceResult](audience), nil
}

func NewCreateAudienceService(
	unitOfWork domain.UnitOfWork,
	logger grpclog.LoggerV2,
) domain.AudienceCreatable {
	return &CreateAudienceService{
		unitOfWork: unitOfWork,
		logger:     logger,
	}
}
