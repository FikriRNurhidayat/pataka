package domain

import (
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
)

type BulkCreateAudiencesAudienceParams struct {
	Audiences []CreateAudienceParams
}

type BulkCreateAudiencesAudienceResult struct {
	Audiences []Audience
}

type CreateAudienceParams struct {
	AudienceId  string
	FeatureName string
	Enabled     bool
}

type CreateAudienceResult struct {
	Audience *Audience
}

type GetAudienceParams struct {
	AudienceId  string
	FeatureName string
}

type GetAudienceResult struct {
	Audience *Audience
}

type ListAudiencesParams struct {
	*pagination.PaginationParams
	Sort        string
	FeatureName string
	AudienceIds []string
	Enabled     *bool
}

type ListAudiencesResult struct {
	*pagination.PaginationResult
	Size      uint32
	Audiences []Audience
}

type UpdateAudienceParams struct {
	AudienceId  string
	FeatureName string
	Enabled     bool
}

type UpdateAudienceResult struct {
	Audience *Audience
}

type DeleteAudienceParams struct {
	AudienceId  string
	FeatureName string
}

type AudienceResultable interface {
	GetAudienceResult | CreateAudienceResult | UpdateAudienceResult
}

func ToAudienceResult[T AudienceResultable](audience *Audience) *T {
	return &T{
		Audience: audience,
	}
}
