package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/manager"
)

func (s *FeatureServer) ListFeatures(ctx context.Context, req *featureflag.ListFeaturesRequest) (*featureflag.ListFeaturesResponse, error) {
	result, err := s.List.Call(ctx, &ListParams{
		PaginationParams: manager.PaginationParams{
			PageNumber: req.GetPageNumber(),
			PageSize:   req.GetPageSize(),
		},
		Sort:    req.GetSort(),
		Q:       req.GetQ(),
		Enabled: req.GetEnabled(),
	})
	if err != nil {
		s.Logger.Infof("[ListFeatures] failed to list a feature resource: %s", err.Error())
		return nil, err
	}

	return ToFeaturesProtoResponse(result), nil
}
