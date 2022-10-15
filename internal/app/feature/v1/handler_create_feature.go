package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) CreateFeature(ctx context.Context, req *featurev1.CreateFeatureRequest) (*featurev1.CreateFeatureResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteFeatureScope); err != nil {
		return nil, err
	}

	result, err := s.createFeatureService.Call(ctx, &domain.CreateFeatureParams{
		Name:    req.GetName(),
		Label:   req.GetLabel(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToFeatureProtoResponse[featurev1.CreateFeatureResponse](result.Feature), nil
}
