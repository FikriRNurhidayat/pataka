package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) GetAudience(ctx context.Context, req *audiencev1.GetAudienceRequest) (*audiencev1.GetAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.ReadAudienceScope); err != nil {
		return nil, err
	}

	result, err := s.getAudienceService.Call(ctx, &domain.GetAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.GetAudienceResponse](result.Audience), nil
}
