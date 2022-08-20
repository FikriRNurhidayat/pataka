package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/manager"
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
	manager.PaginationParams
	Sort    string
	Q       string
	Enabled *bool
}

type ListResult struct {
	manager.PaginationResult
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

type ResultFeatureable interface {
	GetResult | CreateResult | UpdateResult
}

func ToFeatureResult[T ResultFeatureable](feature *Feature) *T {
	return &T{
		Feature: feature,
	}
}
