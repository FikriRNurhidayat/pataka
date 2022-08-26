package domain

import (
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
)

type CreateFeatureParams struct {
	Name    string
	Label   string
	Enabled bool
}

type CreateFeatureResult struct {
	Feature *Feature
}

type ListFeaturesParams struct {
	*pagination.PaginationParams
	Sort    string
	Q       string
	Enabled *bool
}

type ListFeaturesResult struct {
	*pagination.PaginationResult
	Size     uint32
	Features []Feature
}

type GetFeatureParams struct {
	Name string
}

type GetFeatureResult struct {
	Feature *Feature
}

type UpdateFeatureParams struct {
	Name    string
	Label   string
	Enabled bool
}

type UpdateFeatureResult struct {
	Feature *Feature
}

type DeleteFeatureParams struct {
	Name string
}

type FeatureResultable interface {
	GetFeatureResult | CreateFeatureResult | UpdateFeatureResult
}

func ToFeatureResult[T FeatureResultable](feature *Feature) *T {
	return &T{
		Feature: feature,
	}
}
