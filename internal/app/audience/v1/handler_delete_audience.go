package audience

import (
	"context"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) DeleteAudience(ctx context.Context, req *audiencev1.DeleteAudienceRequest) (*audiencev1.DeleteAudienceResponse, error) {
	err := s.Delete.Call(ctx, &DeleteParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return &audiencev1.DeleteAudienceResponse{}, nil
}
