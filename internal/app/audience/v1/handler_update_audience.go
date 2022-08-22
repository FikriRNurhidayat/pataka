package audience

import (
	"context"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) UpdateAudience(ctx context.Context, req *audiencev1.UpdateAudienceRequest) (*audiencev1.UpdateAudienceResponse, error) {
	result, err := s.Update.Call(ctx, &UpdateParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.UpdateAudienceResponse](result.Audience), nil
}
