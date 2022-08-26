package domain

import "context"

type FeatureCreateable interface {
	Call(context.Context, *CreateFeatureParams) (*CreateFeatureResult, error)
}

type FeatureGetable interface {
	Call(context.Context, *GetFeatureParams) (*GetFeatureResult, error)
}

type FeaturesListable interface {
	Call(context.Context, *ListFeaturesParams) (*ListFeaturesResult, error)
}

type FeatureUpdatable interface {
	Call(context.Context, *UpdateFeatureParams) (*UpdateFeatureResult, error)
}

type FeatureDeletable interface {
	Call(context.Context, *DeleteFeatureParams) error
}
