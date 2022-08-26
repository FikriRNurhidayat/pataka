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

type CreateAudienceService struct {
	authentication     auth.Authenticatable
	audienceRepository domain.AudienceRepository
	featureRepository  domain.FeatureRepository
	logger             grpclog.LoggerV2
}

func (s *CreateAudienceService) Call(ctx context.Context, params *domain.CreateAudienceParams) (*domain.CreateAudienceResult, error) {
	if err := s.authentication.Valid(ctx); err != nil {
		return nil, err
	}

	feature, err := s.featureRepository.Get(ctx, params.FeatureName)
	if err != nil {
		s.logger.Error("[create-audience-service] failed to retrieve feature on repository")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return nil, status.Error(codes.NotFound, "Feature not found")
	}

	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Error("[create-audience-service] failed to retrieve audience on repository")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if audience != nil {
		return nil, status.Error(codes.AlreadyExists, "Audience already exist")
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

	if err := s.audienceRepository.Save(ctx, audience); err != nil {
		s.logger.Error("[create-audience-service] failed to save audience on repository")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	feature.HasAudience = true

	if err := s.featureRepository.Save(ctx, feature); err != nil {
		s.logger.Error("[create-audience-service] failed to save feature on repository")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return domain.ToAudienceResult[domain.CreateAudienceResult](audience), nil
}

func NewCreateAudienceService(
	authentication auth.Authenticatable,
	audienceRepository domain.AudienceRepository,
	featureRepository domain.FeatureRepository,
	logger grpclog.LoggerV2,
) domain.AudienceCreatable {
	return &CreateAudienceService{
		authentication:     authentication,
		audienceRepository: audienceRepository,
		featureRepository:  featureRepository,
		logger:             logger,
	}
}
