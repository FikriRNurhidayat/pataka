package audience

import (
	"context"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) GetAudience(ctx context.Context, req *audiencev1.GetAudienceRequest) (*audiencev1.GetAudienceResponse, error) {
	result, err := s.Get.Call(ctx, &GetParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.GetAudienceResponse](result.Audience), nil
}
