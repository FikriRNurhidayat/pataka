package feature

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
)

func (s *FeatureServer) DeleteFeature(ctx context.Context, req *featureflag.DeleteFeatureRequest) (*featureflag.DeleteFeatureResponse, error)
