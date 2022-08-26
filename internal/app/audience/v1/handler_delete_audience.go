package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) DeleteAudience(ctx context.Context, req *audiencev1.DeleteAudienceRequest) (*audiencev1.DeleteAudienceResponse, error) {
	err := s.deleteAudienceService.Call(ctx, &domain.DeleteAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return &audiencev1.DeleteAudienceResponse{}, nil
}
