package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type DeleteAudienceService struct {
	authentication     auth.Authenticatable
	audienceRepository AudienceRepository
	logger             grpclog.LoggerV2
}

func (s *DeleteAudienceService) Call(ctx context.Context, params *DeleteParams) error {
	if err := s.authentication.Valid(ctx); err != nil {
		return err
	}

	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Errorf("[delete-audience-service] failed to retrieve a audience resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	if audience == nil {
		return status.Error(codes.NotFound, "Audience not found")
	}

	if err := s.audienceRepository.Delete(ctx, audience); err != nil {
		s.logger.Errorf("[delete-audience-service] failed to delete a audience resource: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}

func NewDeleteAudienceService(
	authentication auth.Authenticatable,
	audienceRepository AudienceRepository,
	logger grpclog.LoggerV2,
) Deletable {
	return &DeleteAudienceService{
		authentication:     authentication,
		audienceRepository: audienceRepository,
		logger:             logger,
	}
}
