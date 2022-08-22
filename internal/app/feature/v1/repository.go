package feature

import "context"

type FeatureRepository interface {
	Save(ctx context.Context, feature *Feature) error
	Get(ctx context.Context, name string) (*Feature, error)
	Delete(ctx context.Context, name string) error
	List(ctx context.Context, args *FeatureListArgs) ([]Feature, error)
	Size(ctx context.Context, args *FeatureFilterArgs) (uint32, error)
}

type FeatureFilterArgs struct {
	Q       string `db:"q"`
	Name    string `db:"name"`
	Label   string `db:"label"`
	Enabled *bool  `db:"enabled"`
}

type FeatureListArgs struct {
	Limit  uint32 `db:"limit"`
	Offset uint32 `db:"offset"`
	Sort   string
	Filter *FeatureFilterArgs
}