package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *Server) GetFeature(ctx context.Context, req *featureflagpb.GetFeatureRequest) (*featureflagpb.GetFeatureResponse, error) {
	result, err := s.Get.Call(ctx, &GetParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Info("[get-feature-handler] failed to get a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featureflagpb.GetFeatureResponse](result.Feature), nil
}
