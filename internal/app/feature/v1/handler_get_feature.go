package feature

import (
	"context"

	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) GetFeature(ctx context.Context, req *featurev1.GetFeatureRequest) (*featurev1.GetFeatureResponse, error) {
	result, err := s.Get.Call(ctx, &GetParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Info("[get-feature-handler] failed to get a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.GetFeatureResponse](result.Feature), nil
}
