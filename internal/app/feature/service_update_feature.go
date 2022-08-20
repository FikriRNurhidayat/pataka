package feature

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type Updatable interface {
	Call(context.Context, *UpdateParams) (*UpdateResult, error)
}

type UpdateFeatureService struct {
	Authentication    auth.Authenticatable
	FeatureRepository FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *UpdateFeatureService) Call(ctx context.Context, params *UpdateParams) (*UpdateResult, error) {
	if err := s.Authentication.Valid(ctx); err != nil {
		return nil, err
	}

	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[update-feature-service] failed to retrieve a feature resource: %s", err.Error())
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

	if err := s.FeatureRepository.Save(ctx, feature); err != nil {
		s.Logger.Errorf("[update-feature-service] failed to save a feature resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return ToFeatureResult[UpdateResult](feature), nil
}

func NewUpdateFeatureService(
	Authentication auth.Authenticatable,
	FeatureRepository FeatureRepository,
	Logger grpclog.LoggerV2,
) Updatable {
	return &UpdateFeatureService{
		Authentication:    Authentication,
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
