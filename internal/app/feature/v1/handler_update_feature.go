package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) UpdateFeature(ctx context.Context, req *featurev1.UpdateFeatureRequest) (*featurev1.UpdateFeatureResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteFeatureScope); err != nil {
		return nil, err
	}

	result, err := s.updateFeatureService.Call(ctx, &domain.UpdateFeatureParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.logger.Info("[update-feature-handler] failed to update a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.UpdateFeatureResponse](result.Feature), nil
}
