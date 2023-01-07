package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
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

func (s *Server) ListFeatures(ctx context.Context, req *featurev1.ListFeaturesRequest) (*featurev1.ListFeaturesResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.ReadFeatureScope); err != nil {
		return nil, err
	}

	var (
		enabled *bool
		e       bool
	)
	switch req.GetStatus() {
	case featurev1.FeatureStatus_FEATURE_STATUS_UNSPECIFIED:
		enabled = nil
	case featurev1.FeatureStatus_FEATURE_STATUS_ENABLED:
		e = true
		enabled = &e
	case featurev1.FeatureStatus_FEATURE_STATUS_DISABLED:
		e = false
		enabled = &e
	}

	result, err := s.listFeaturesService.Call(ctx, &domain.ListFeaturesParams{
		PaginationParams: &pagination.PaginationParams{
			PageNumber: req.GetPageNumber(),
			PageSize:   req.GetPageSize(),
		},
		Sort:    req.GetSort(),
		Q:       req.GetQ(),
		Enabled: enabled,
	})

	if err != nil {
		s.logger.Info("[list-features-handler] failed to list feature collection")
		return nil, err
	}

	return ToFeaturesProtoResponse(result), nil
}

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

func (s *Server) DeleteFeature(ctx context.Context, req *featurev1.DeleteFeatureRequest) (*featurev1.DeleteFeatureResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteFeatureScope); err != nil {
		return nil, err
	}

	err := s.deleteFeatureService.Call(ctx, &domain.DeleteFeatureParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.logger.Info("[delete-feature-handler] failed to delete a feature resource")
		return nil, err
	}

	return &featurev1.DeleteFeatureResponse{}, nil
}
