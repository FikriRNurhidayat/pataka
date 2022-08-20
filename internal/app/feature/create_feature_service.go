package feature

import (
	"context"
	"errors"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/grpclog"
)

type Createable interface {
	Call(context.Context, *CreateParams) (*CreateResult, error)
}

type CreateFeatureService struct {
	FeatureRepository domain.FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *CreateFeatureService) Call(ctx context.Context, params *CreateParams) (*CreateResult, error) {
	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to retrieve a feature resource: %v", err.Error())
		return nil, err
	}

	if feature != nil {
		return nil, errors.New("Feature already exists")
	}

	feature = &domain.Feature{
		Name:      params.Name,
		Label:     params.Label,
		Enabled:   params.Enabled,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if params.Enabled {
		feature.EnabledAt = time.Now().Local()
	}

	if err := s.FeatureRepository.Save(ctx, feature); err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to save a feature resource: %v", err.Error())
		return nil, err
	}

	return &CreateResult{feature}, nil
}

func NewCreateFeatureService(
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Createable {
	return &CreateFeatureService{
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
