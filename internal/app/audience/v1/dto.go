package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
)

type BulkCreateParams struct {
	Audiences []CreateParams
}

type BulkCreateResult struct {
	Audiences []Audience
}

type CreateParams struct {
	AudienceId  string
	FeatureName string
	Enabled     bool
}

type CreateResult struct {
	Audience *Audience
}

type GetParams struct {
	AudienceId  string
	FeatureName string
}

type GetResult struct {
	Audience *Audience
}

type ListParams struct {
	*pagination.PaginationParams
	Sort        string
	FeatureName string
	AudienceIds []string
	Enabled     *bool
}

type ListResult struct {
	*pagination.PaginationResult
	Size      uint32
	Audiences []Audience
}

type UpdateParams struct {
	AudienceId  string
	FeatureName string
	Enabled     bool
}

type UpdateResult struct {
	Audience *Audience
}

type DeleteParams struct {
	AudienceId  string
	FeatureName string
}

type AudienceResultable interface {
	GetResult | CreateResult | UpdateResult
}

func ToAudienceResult[T AudienceResultable](audience *Audience) *T {
	return &T{
		Audience: audience,
	}
}
