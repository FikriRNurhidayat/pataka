package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
)

type CreateParams struct {
	Name    string
	Label   string
	Enabled bool
}

type CreateResult struct {
	Feature *Feature
}

type ListParams struct {
	*pagination.PaginationParams
	Sort    string
	Q       string
	Enabled *bool
}

type ListResult struct {
	*pagination.PaginationResult
	Size     uint32
	Features []Feature
}

type GetParams struct {
	Name string
}

type GetResult struct {
	Feature *Feature
}

type UpdateParams struct {
	Name    string
	Label   string
	Enabled bool
}

type UpdateResult struct {
	Feature *Feature
}

type DeleteParams struct {
	Name string
}

type FeatureResultable interface {
	GetResult | CreateResult | UpdateResult
}

func ToFeatureResult[T FeatureResultable](feature *Feature) *T {
	return &T{
		Feature: feature,
	}
}
