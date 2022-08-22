package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *Server) CreateFeature(ctx context.Context, req *featureflagpb.CreateFeatureRequest) (*featureflagpb.CreateFeatureResponse, error) {
	result, err := s.Create.Call(ctx, &CreateParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.Logger.Info("[create-feature-handler] failed to create a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featureflagpb.CreateFeatureResponse](result.Feature), nil
}
