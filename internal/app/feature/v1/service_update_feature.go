package feature

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type UpdateFeatureService struct {
	authentication    auth.Authenticatable
	featureRepository domain.FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *UpdateFeatureService) Call(ctx context.Context, params *domain.UpdateFeatureParams) (*domain.UpdateFeatureResult, error) {
	if err := s.authentication.Valid(ctx); err != nil {
		return nil, err
	}

	feature, err := s.featureRepository.Get(ctx, params.Name)
	if err != nil {
		s.logger.Errorf("[update-feature-service] failed to retrieve a feature resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return nil, status.Error(codes.NotFound, "Feature not found")
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

	if err := s.featureRepository.Save(ctx, feature); err != nil {
		s.logger.Errorf("[update-feature-service] failed to save a feature resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return domain.ToFeatureResult[domain.UpdateFeatureResult](feature), nil
}

func NewUpdateFeatureService(
	authentication auth.Authenticatable,
	featureRepository domain.FeatureRepository,
	logger grpclog.LoggerV2,
) domain.FeatureUpdatable {
	return &UpdateFeatureService{
		authentication:    authentication,
		featureRepository: featureRepository,
		logger:            logger,
	}
}
