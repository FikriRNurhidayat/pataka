package feature

import (
	"context"
	"database/sql"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type Updatable interface {
	Call(context.Context, *UpdateParams) (*UpdateResult, error)
}

type UpdateFeatureService struct {
	AuthenticationService authentication.Authenticatable
	FeatureRepository     domain.FeatureRepository
	Logger                grpclog.LoggerV2
}

func (s *UpdateFeatureService) Call(ctx context.Context, params *UpdateParams) (*UpdateResult, error) {
	if err := s.AuthenticationService.Valid(ctx); err != nil {
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
		feature.EnabledAt = sql.NullTime{
			Time:  time.Now().Local(),
			Valid: true,
		}
	}

	// Update EnabledAt by zero
	// If the feature is initially enabled then disabled by this action
	if feature.Enabled && !params.Enabled {
		feature.EnabledAt = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
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
	AuthenticationService authentication.Authenticatable,
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Updatable {
	return &UpdateFeatureService{
		AuthenticationService: AuthenticationService,
		FeatureRepository:     FeatureRepository,
		Logger:                Logger,
	}
}
