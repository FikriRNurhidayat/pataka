package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) ListFeatures(ctx context.Context, req *featurev1.ListFeaturesRequest) (*featurev1.ListFeaturesResponse, error) {
	// TODO: Find a better way
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
		s.Logger.Info("[list-features-handler] failed to list feature collection")
		return nil, err
	}

	return ToFeaturesProtoResponse(result), nil
}
