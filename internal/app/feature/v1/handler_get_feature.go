package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) GetFeature(ctx context.Context, req *featurev1.GetFeatureRequest) (*featurev1.GetFeatureResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.ReadFeatureScope); err != nil {
		return nil, err
	}

	result, err := s.getFeatureService.Call(ctx, &domain.GetFeatureParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.logger.Info("[get-feature-handler] failed to get a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.GetFeatureResponse](result.Feature), nil
}
