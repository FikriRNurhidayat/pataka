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

type CreateFeatureService struct {
	authentication    auth.Authenticatable
	featureRepository domain.FeatureRepository
	logger            grpclog.LoggerV2
}

func (s *CreateFeatureService) Call(ctx context.Context, params *domain.CreateFeatureParams) (*domain.CreateFeatureResult, error) {
	if err := s.authentication.Valid(ctx); err != nil {
		return nil, err
	}

	feature, err := s.featureRepository.Get(ctx, params.Name)
	if err != nil {
		s.logger.Error("[create-feature-service] failed to retrieve a feature resource")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if feature != nil {
		return nil, status.Error(codes.InvalidArgument, "Feature already exists")
	}

	feature = &domain.Feature{
		Name:             params.Name,
		Label:            params.Label,
		Enabled:          params.Enabled,
		HasAudience:      false,
		HasAudienceGroup: false,
		CreatedAt:        time.Now().Local(),
		UpdatedAt:        time.Now().Local(),
		EnabledAt:        time.Time{},
	}

	if params.Enabled {
		feature.EnabledAt = time.Now().Local()
	}

	if err := s.featureRepository.Save(ctx, feature); err != nil {
		s.logger.Error("[create-feature-service] failed to save a feature resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return domain.ToFeatureResult[domain.CreateFeatureResult](feature), nil
}

func NewCreateFeatureService(
	authentication auth.Authenticatable,
	featureRepository domain.FeatureRepository,
	logger grpclog.LoggerV2,
) domain.FeatureCreateable {
	return &CreateFeatureService{
		authentication:    authentication,
		featureRepository: featureRepository,
		logger:            logger,
	}
}
