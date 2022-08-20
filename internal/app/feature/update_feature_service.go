package feature

import (
	"context"
	"errors"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/grpclog"
)

type Updatable interface {
	Call(context.Context, *UpdateParams) (*UpdateResult, error)
}

type UpdateFeatureService struct {
	FeatureRepository domain.FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *UpdateFeatureService) Call(ctx context.Context, params *UpdateParams) (*UpdateResult, error) {
	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to retrieve a feature resource: %v", err.Error())
		return nil, err
	}

	if feature == nil {
		return nil, errors.New("Feature not found")
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

	if err := s.FeatureRepository.Save(ctx, feature); err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to save a feature resource: %v", err.Error())
		return nil, err
	}

	return &UpdateResult{feature}, nil
}

func NewUpdateFeatureService(
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Updatable {
	return &UpdateFeatureService{
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
