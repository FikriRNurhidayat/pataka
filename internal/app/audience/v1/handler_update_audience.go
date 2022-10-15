package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) UpdateAudience(ctx context.Context, req *audiencev1.UpdateAudienceRequest) (*audiencev1.UpdateAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteAudienceScope); err != nil {
		return nil, err
	}

	result, err := s.updateAudienceService.Call(ctx, &domain.UpdateAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.UpdateAudienceResponse](result.Audience), nil
}
