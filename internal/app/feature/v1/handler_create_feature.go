package feature

import (
	"context"

	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) CreateFeature(ctx context.Context, req *featurev1.CreateFeatureRequest) (*featurev1.CreateFeatureResponse, error) {
	result, err := s.Create.Call(ctx, &CreateParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.Logger.Info("[create-feature-handler] failed to create a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.CreateFeatureResponse](result.Feature), nil
}
