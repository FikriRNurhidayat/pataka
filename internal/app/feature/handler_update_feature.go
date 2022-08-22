package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *Server) UpdateFeature(ctx context.Context, req *featureflagpb.UpdateFeatureRequest) (*featureflagpb.UpdateFeatureResponse, error) {
	result, err := s.Update.Call(ctx, &UpdateParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.Logger.Info("[update-feature-handler] failed to update a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featureflagpb.UpdateFeatureResponse](result.Feature), nil
}
