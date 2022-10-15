package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type DeleteFeatureService struct {
	unitOfWork domain.UnitOfWork
	logger     grpclog.LoggerV2
}

func (s *DeleteFeatureService) Call(ctx context.Context, params *domain.DeleteFeatureParams) error {
	return s.unitOfWork.Do(ctx, func(r domain.Repository) error {
		feature, err := r.FeatureRepository().Get(ctx, params.Name)
		if err != nil {
			s.logger.Errorf("[delete-feature-service] failed to retrieve a feature resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		if feature == nil {
			return status.Error(codes.NotFound, "Feature not found")
		}

		if err := r.FeatureRepository().Delete(ctx, feature.Name); err != nil {
			s.logger.Errorf("[delete-feature-service] failed to delete a feature resource: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		if err := r.AudienceRepository().DeleteBy(ctx, &domain.AudienceFilterArgs{
			FeatureName: feature.Name,
		}); err != nil {
			s.logger.Errorf("[delete-feature-service] failed to bulk delete audience collections: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		return nil
	})
}

func NewDeleteFeatureService(unitOfWork domain.UnitOfWork, logger grpclog.LoggerV2) domain.FeatureDeletable {
	return &DeleteFeatureService{
		unitOfWork: unitOfWork,
		logger:     logger,
	}
}
