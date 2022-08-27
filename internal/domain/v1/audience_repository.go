package domain

import (
	"context"
)

type AudienceRepository interface {
	Get(ctx context.Context, featureName string, audienceId string) (*Audience, error)
	GetBy(ctx context.Context, args *AudienceGetByArgs) (*Audience, error)
	List(context.Context, *AudienceListArgs) ([]Audience, error)
	Size(context.Context, *AudienceFilterArgs) (uint32, error)
	Save(context.Context, *Audience) error
	Delete(context.Context, *Audience) error
	DeleteBy(ctx context.Context, args *AudienceFilterArgs) error
}

type AudienceFilterArgs struct {
	FeatureName string   `db:"feature_name"`
	AudienceIds []string `db:"audience_id"`
	Enabled     *bool    `db:"enabled"`
}

type AudienceGetByArgs struct {
	Sort   string
	Filter *AudienceFilterArgs
}

type AudienceListArgs struct {
	Limit  uint32 `db:"limit"`
	Offset uint32 `db:"offset"`
	Sort   string
	Filter *AudienceFilterArgs
}
