package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) CreateAudience(ctx context.Context, req *audiencev1.CreateAudienceRequest) (*audiencev1.CreateAudienceResponse, error) {
	result, err := s.createAudienceService.Call(ctx, &domain.CreateAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.CreateAudienceResponse](result.Audience), nil
}
