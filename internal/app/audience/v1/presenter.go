package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

type AudienceProtoable interface {
	audiencev1.GetAudienceResponse | audiencev1.UpdateAudienceResponse | audiencev1.CreateAudienceResponse
}

func ToAudiencesProtoResponse(result *domain.ListAudiencesResult) *audiencev1.ListAudiencesResponse {
	res := &audiencev1.ListAudiencesResponse{
		PageNumber: result.PageNumber,
		PageSize:   result.PageSize,
		PageCount:  result.PageCount,
		Size:       result.Size,
	}

	for _, audience := range result.Audiences {
		res.Audiences = append(res.Audiences, ToAudienceProto(&audience))
	}

	return res
}

func ToAudienceProtoResponse[T AudienceProtoable](audience *domain.Audience) *T {
	return &T{
		Audience: ToAudienceProto(audience),
	}
}

func ToAudienceProto(audience *domain.Audience) *audiencev1.Audience {
	res := &audiencev1.Audience{
		AudienceId:  audience.AudienceId,
		FeatureName: audience.FeatureName,
		Enabled:     audience.Enabled,
		CreatedAt:   timestamppb.New(audience.CreatedAt),
		UpdatedAt:   timestamppb.New(audience.UpdatedAt),
	}

	if !audience.EnabledAt.IsZero() {
		res.EnabledAt = timestamppb.New(audience.EnabledAt)
	}

	return res
}
