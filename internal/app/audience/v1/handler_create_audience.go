package audience

import (
	"context"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) CreateAudience(ctx context.Context, req *audiencev1.CreateAudienceRequest) (*audiencev1.CreateAudienceResponse, error) {
	result, err := s.Create.Call(ctx, &CreateParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.CreateAudienceResponse](result.Audience), nil
}
