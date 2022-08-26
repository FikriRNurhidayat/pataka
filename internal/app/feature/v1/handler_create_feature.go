package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) CreateFeature(ctx context.Context, req *featurev1.CreateFeatureRequest) (*featurev1.CreateFeatureResponse, error) {
	result, err := s.createFeatureService.Call(ctx, &domain.CreateFeatureParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.logger.Info("[create-feature-handler] failed to create a feature resource")
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.CreateFeatureResponse](result.Feature), nil
}
