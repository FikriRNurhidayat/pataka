package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

func (s *Server) DeleteFeature(ctx context.Context, req *featurev1.DeleteFeatureRequest) (*featurev1.DeleteFeatureResponse, error) {
	err := s.deleteFeatureService.Call(ctx, &domain.DeleteFeatureParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Info("[delete-feature-handler] failed to delete a feature resource")
		return nil, err
	}

	return &featurev1.DeleteFeatureResponse{}, nil
}
