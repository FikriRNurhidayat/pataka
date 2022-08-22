package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/manager"
)

func (s *Server) ListFeatures(ctx context.Context, req *featureflagpb.ListFeaturesRequest) (*featureflagpb.ListFeaturesResponse, error) {
	// TODO: Find a better way
	var (
		enabled *bool
		e       bool
	)
	switch req.GetStatus() {
	case featureflagpb.FeatureStatus_FEATURE_STATUS_UNSPECIFIED:
		enabled = nil
	case featureflagpb.FeatureStatus_FEATURE_STATUS_ENABLED:
		e = true
		enabled = &e
	case featureflagpb.FeatureStatus_FEATURE_STATUS_DISABLED:
		e = false
		enabled = &e
	}

	result, err := s.List.Call(ctx, &ListParams{
		PaginationParams: manager.PaginationParams{
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
