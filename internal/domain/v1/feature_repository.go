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
	Q       string `db:"q" url:"q"`
	Name    string `db:"name" url:"name"`
	Label   string `db:"label" url:"label"`
	Enabled *bool  `db:"enabled" url:"enabled"`
}

type FeatureGetByArgs struct {
	Sort   string             `url:"sort"`
	Filter *FeatureFilterArgs `url:"filter"`
}

type FeatureListArgs struct {
	Limit  uint32             `db:"limit" url:"limit"`
	Offset uint32             `db:"offset" url:"offset"`
	Sort   string             `url:"sort"`
	Filter *FeatureFilterArgs `url:"filter"`
}
