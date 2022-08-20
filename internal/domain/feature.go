package domain

import (
	"context"
	"database/sql"
	"time"
)

type Feature struct {
	Name             string       `json:"name" db:"name"`
	Label            string       `json:"label" db:"label"`
	Enabled          bool         `json:"enabled" db:"enabled"`
	HasAudience      bool         `json:"has_audience" db:"has_audience"`
	HasAudienceGroup bool         `json:"has_audience_group" db:"has_audience_group"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" db:"updated_at"`
	EnabledAt        sql.NullTime `json:"enabled_at" db:"enabled_at"`
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
