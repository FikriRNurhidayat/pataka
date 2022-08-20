package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
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
		s.Logger.Error("[get-feature-service] failed to retrieve a feature resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return nil, status.Error(codes.NotFound, "Feature not found")
	}

	return ToFeatureResult[GetResult](feature), nil
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
