package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type DeleteAudienceService struct {
	unitOfWork domain.UnitOfWork
	logger     grpclog.LoggerV2
}

func (s *DeleteAudienceService) Call(ctx context.Context, params *domain.DeleteAudienceParams) error {
	return s.unitOfWork.Do(ctx, func(r domain.Repository) error {
		audience, err := r.AudienceRepository().Get(ctx, params.FeatureName, params.AudienceId)
		if err != nil {
			s.logger.Error("[delete-audience-service] failed to retrieve a audience on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		if audience == nil {
			return status.Error(codes.NotFound, "Audience not found")
		}

		if err := r.AudienceRepository().Delete(ctx, audience); err != nil {
			s.logger.Error("[delete-audience-service] failed to delete a audience on repository")
			return status.Error(codes.Internal, "Internal server error")
		}

		return nil
	})
}

func NewDeleteAudienceService(
	unitOfWork domain.UnitOfWork,
	logger grpclog.LoggerV2,
) domain.AudienceDeletable {
	return &DeleteAudienceService{
		unitOfWork: unitOfWork,
		logger:     logger,
	}
}
