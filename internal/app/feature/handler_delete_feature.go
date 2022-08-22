package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *Server) DeleteFeature(ctx context.Context, req *featureflagpb.DeleteFeatureRequest) (*featureflagpb.DeleteFeatureResponse, error) {
	err := s.Delete.Call(ctx, &DeleteParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Info("[delete-feature-handler] failed to delete a feature resource")
		return nil, err
	}

	return &featureflagpb.DeleteFeatureResponse{}, nil
}
