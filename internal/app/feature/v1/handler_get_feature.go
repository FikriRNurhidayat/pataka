package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) GetFeature(ctx context.Context, req *featurev1.GetFeatureRequest) (*featurev1.GetFeatureResponse, error) {
	result, err := s.getFeatureService.Call(ctx, &domain.GetFeatureParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Info("[get-feature-handler] failed to get a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.GetFeatureResponse](result.Feature), nil
}
