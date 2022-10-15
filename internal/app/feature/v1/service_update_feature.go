package feature

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type UpdateFeatureService struct {
	unitOfWork domain.UnitOfWork
	logger     grpclog.LoggerV2
}

func (s *UpdateFeatureService) Call(ctx context.Context, params *domain.UpdateFeatureParams) (*domain.UpdateFeatureResult, error) {
	var feature *domain.Feature

	if err := s.unitOfWork.Do(ctx, func(r domain.Repository) (err error) {
		feature, err = r.FeatureRepository().Get(ctx, params.Name)
		if err != nil {
			s.logger.Errorf("[update-feature-service] failed to retrieve a feature resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		if feature == nil {
			return status.Error(codes.NotFound, "Feature not found")
		}

		// Update EnabledAt by now
		// If the feature is initially disabled then enabled by this action
		if !feature.Enabled && params.Enabled {
			feature.EnabledAt = time.Now().Local()
		}

		// Update EnabledAt by zero
		// If the feature is initially enabled then disabled by this action
		if feature.Enabled && !params.Enabled {
			feature.EnabledAt = time.Time{}
		}

		feature.Label = params.Label
		feature.Enabled = params.Enabled

		if err := r.FeatureRepository().Save(ctx, feature); err != nil {
			s.logger.Errorf("[update-feature-service] failed to save a feature resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return domain.ToFeatureResult[domain.UpdateFeatureResult](feature), nil
}

func NewUpdateFeatureService(unitOfWork domain.UnitOfWork, logger grpclog.LoggerV2) domain.FeatureUpdatable {
	return &UpdateFeatureService{
		unitOfWork: unitOfWork,
		logger:     logger,
	}
}
