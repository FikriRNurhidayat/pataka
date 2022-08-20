package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *FeatureServer) CreateFeature(ctx context.Context, req *featureflag.CreateFeatureRequest) (*featureflag.CreateFeatureResponse, error) {
	result, err := s.Create.Call(ctx, &CreateParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.Logger.Infof("[create-feature-handler] failed to create a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featureflag.CreateFeatureResponse](result.Feature), nil
}
