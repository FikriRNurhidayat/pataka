package feature

import (
	"context"
	"errors"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/grpclog"
)

type Getable interface {
	Call(context.Context, *GetParams) (*GetResult, error)
}

type GetFeatureService struct {
	FeatureRepository domain.FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *GetFeatureService) Call(ctx context.Context, params *GetParams) (*GetResult, error) {
	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to retrieve a feature resource: %s", err.Error())
		return nil, err
	}

	if feature == nil {
		return nil, errors.New("Feature not found")
	}

	return &GetResult{feature}, nil
}

func NewGetFeatureService(
	FeatureRepository domain.FeatureRepository,
	Logger grpclog.LoggerV2,
) Getable {
	return &GetFeatureService{
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
