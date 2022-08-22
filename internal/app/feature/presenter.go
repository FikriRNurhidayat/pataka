package feature

import (
	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProtoFeatureable interface {
	featureflagpb.GetFeatureResponse | featureflagpb.UpdateFeatureResponse | featureflagpb.CreateFeatureResponse
}

func ToFeaturesProtoResponse(result *ListResult) *featureflagpb.ListFeaturesResponse {
	res := &featureflagpb.ListFeaturesResponse{
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

func ToFeatureProto(feature *Feature) *featureflagpb.Feature {
	res := &featureflagpb.Feature{
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
