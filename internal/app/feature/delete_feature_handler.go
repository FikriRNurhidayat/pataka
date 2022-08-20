package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *FeatureServer) DeleteFeature(ctx context.Context, req *featureflag.DeleteFeatureRequest) (*featureflag.DeleteFeatureResponse, error) {
	err := s.Delete.Call(ctx, &DeleteParams{
		Name: req.GetName(),
	})
	if err != nil {
		s.Logger.Infof("[DeleteFeature] failed to delete a feature resource: %s", err.Error())
		return nil, err
	}

	return &featureflag.DeleteFeatureResponse{}, nil
}
