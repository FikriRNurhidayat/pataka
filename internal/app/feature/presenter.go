package feature

import (
	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProtoFeatureable interface {
	featureflag.GetFeatureResponse | featureflag.UpdateFeatureResponse | featureflag.CreateFeatureResponse
}

func ToFeaturesProtoResponse(result *ListResult) *featureflag.ListFeaturesResponse {
	res := &featureflag.ListFeaturesResponse{
		PageNumber: result.PageNumber,
		PageSize:   result.PageSize,
		PageCount:  result.PageCount,
		Size:       result.Size,
	}

	for _, feature := range result.Features {
		res.Features = append(res.Features, ToFeatureProto(&feature))
	}

	return res
}

func ToFeatureProtoResponse[T ProtoFeatureable](feature *domain.Feature) *T {
	return &T{
		Feature: ToFeatureProto(feature),
	}
}

func ToFeatureProto(feature *domain.Feature) *featureflag.Feature {
	res := &featureflag.Feature{
		Name:      feature.Name,
		Label:     feature.Label,
		Enabled:   feature.Enabled,
		CreatedAt: timestamppb.New(feature.CreatedAt),
		UpdatedAt: timestamppb.New(feature.UpdatedAt),
	}

	if !feature.EnabledAt.Time.IsZero() {
		res.EnabledAt = timestamppb.New(feature.EnabledAt.Time)
	}

	return res
}
