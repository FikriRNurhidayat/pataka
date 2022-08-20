package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *Server) GetFeature(ctx context.Context, req *featureflag.GetFeatureRequest) (*featureflag.GetFeatureResponse, error) {
	result, err := s.Get.Call(ctx, &GetParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Infof("[GetFeature] failed to get a feature resource: %s", err.Error())
		return nil, err
	}

	return ToFeatureProtoResponse[featureflag.GetFeatureResponse](result.Feature), nil
}
