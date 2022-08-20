package domain

import (
	"context"
	"time"
)

type Feature struct {
	Name      string    `json:"name"`
	Label     string    `json:"label"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	EnabledAt time.Time `json:"enabled_at"`
}

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
	Enabled bool   `db:"enabled"`
}

type FeatureListArgs struct {
	Limit  uint32 `db:"limit"`
	Offset uint32 `db:"offset"`
	Sort   string
	Filter *FeatureFilterArgs
}
