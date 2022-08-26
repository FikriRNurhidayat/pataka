package domain

import (
	"context"
)

type FeatureRepository interface {
	Save(ctx context.Context, feature *Feature) error
	Get(ctx context.Context, name string) (*Feature, error)
	GetBy(ctx context.Context, args *FeatureGetByArgs) (*Feature, error)
	List(ctx context.Context, args *FeatureListArgs) ([]Feature, error)
	Delete(ctx context.Context, name string) error
	Size(ctx context.Context, args *FeatureFilterArgs) (uint32, error)
	DeleteBy(ctx context.Context, args *FeatureFilterArgs) error
}

type FeatureFilterArgs struct {
	Q       string `db:"q"`
	Name    string `db:"name"`
	Label   string `db:"label"`
	Enabled *bool  `db:"enabled"`
}

type FeatureGetByArgs struct {
	Sort   string
	Filter *FeatureFilterArgs
}

type FeatureListArgs struct {
	Limit  uint32 `db:"limit"`
	Offset uint32 `db:"offset"`
	Sort   string
	Filter *FeatureFilterArgs
}
