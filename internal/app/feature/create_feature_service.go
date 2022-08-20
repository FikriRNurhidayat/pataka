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

type Createable interface {
	Call(context.Context, *CreateParams) (*CreateResult, error)
}

type CreateFeatureService struct {
	AuthenticationService authentication.Authenticatable
	FeatureRepository     domain.FeatureRepository
	Logger                grpclog.LoggerV2
}

func (s *CreateFeatureService) Call(ctx context.Context, params *CreateParams) (*CreateResult, error) {
	if err := s.AuthenticationService.Valid(ctx); err != nil {
		return nil, err
	}

	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Error("[create-feature-service] failed to retrieve a feature resource")
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
		EnabledAt:        sql.NullTime{},
	}

	if params.Enabled {
		feature.EnabledAt.Time = time.Now().Local()
		feature.EnabledAt.Valid = true
	}

	if err := s.FeatureRepository.Save(ctx, feature); err != nil {
		s.Logger.Error("[create-feature-service] failed to save a feature resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return ToFeatureResult[CreateResult](feature), nil
}

func NewCreateFeatureService(
	AuthenticationService authentication.Authenticatable,
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Createable {
	return &CreateFeatureService{
		AuthenticationService: AuthenticationService,
		FeatureRepository:     FeatureRepository,
		Logger:                Logger,
	}
}
