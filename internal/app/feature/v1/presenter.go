package feature

import (
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProtoFeatureable interface {
	featurev1.GetFeatureResponse | featurev1.UpdateFeatureResponse | featurev1.CreateFeatureResponse
}

func ToFeaturesProtoResponse(result *ListResult) *featurev1.ListFeaturesResponse {
	res := &featurev1.ListFeaturesResponse{
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

func ToFeatureProtoResponse[T ProtoFeatureable](feature *Feature) *T {
	return &T{
		Feature: ToFeatureProto(feature),
	}
}

func ToFeatureProto(feature *Feature) *featurev1.Feature {
	res := &featurev1.Feature{
		Name:      feature.Name,
		Label:     feature.Label,
		Enabled:   feature.Enabled,
		CreatedAt: timestamppb.New(feature.CreatedAt),
		UpdatedAt: timestamppb.New(feature.UpdatedAt),
	}

	if !feature.EnabledAt.IsZero() {
		res.EnabledAt = timestamppb.New(feature.EnabledAt)
	}

	return res
}
