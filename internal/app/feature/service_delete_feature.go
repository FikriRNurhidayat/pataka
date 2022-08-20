package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type Deletable interface {
	Call(context.Context, *DeleteParams) error
}

type DeleteFeatureService struct {
	Authentication    auth.Authenticatable
	FeatureRepository FeatureRepository
	Logger            grpclog.LoggerV2
}

func (s *DeleteFeatureService) Call(ctx context.Context, params *DeleteParams) error {
	if err := s.Authentication.Valid(ctx); err != nil {
		return err
	}

	feature, err := s.FeatureRepository.Get(ctx, params.Name)
	if err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to retrieve a feature resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	if feature == nil {
		return status.Error(codes.NotFound, "Feature not found")
	}

	if err := s.FeatureRepository.Delete(ctx, feature.Name); err != nil {
		s.Logger.Errorf("[FeatureRepository] failed to delete a feature resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}

func NewDeleteFeatureService(
	Authentication auth.Authenticatable,
	FeatureRepository FeatureRepository,
	Logger grpclog.LoggerV2,
) Deletable {
	return &DeleteFeatureService{
		Authentication:    Authentication,
		FeatureRepository: FeatureRepository,
		Logger:            Logger,
	}
}
