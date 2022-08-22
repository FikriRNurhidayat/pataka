package audience

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type CreateAudienceService struct {
	authentication     auth.Authenticatable
	audienceRepository AudienceRepository
	logger             grpclog.LoggerV2
}

func (s *CreateAudienceService) Call(ctx context.Context, params *CreateParams) (*CreateResult, error) {
	if err := s.authentication.Valid(ctx); err != nil {
		return nil, err
	}

	audience, err := s.audienceRepository.Get(ctx, params.FeatureName, params.AudienceId)
	if err != nil {
		s.logger.Errorf("[enable-audience-service] failed to retrieve a audience resource: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if audience != nil {
		return nil, status.Error(codes.AlreadyExists, "Audience already exist")
	}

	audience = &Audience{
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
		s.logger.Error("[create-audience-service] failed to save a audience resource")
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return ToAudienceResult[CreateResult](audience), nil
}

func NewCreateAudienceService(
	authentication auth.Authenticatable,
	audienceRepository AudienceRepository,
	logger grpclog.LoggerV2,
) Createable {
	return &CreateAudienceService{
		authentication:     authentication,
		audienceRepository: audienceRepository,
		logger:             logger,
	}
}
