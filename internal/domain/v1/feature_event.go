package domain

import "context"

const (
	FEATURE_CREATED_EVENT = "features.created"
	FEATURE_UPDATED_EVENT = "features.updated"
	FEATURE_GOTTEN_EVENT  = "features.gotten"
	FEATURE_DELETED_EVENT = "features.deleted"
	FEATURE_LISTED_EVENT  = "features.listed"
)

type FeatureCreatedEvent Event[Feature]
type FeatureUpdatedEvent Event[Feature]
type FeatureGottenEvent Event[Feature]
type FeatureListedEvent Event[[]Feature]
type FeatureDeletedEvent Event[Feature]

type FeatureEventEmitter interface {
	EmitCreatedEvent(context.Context, *FeatureCreatedEvent) error
	EmitUpdatedEvent(context.Context, *FeatureUpdatedEvent) error
	EmitDeletedEvent(context.Context, *FeatureDeletedEvent) error
	EmitGottenEvent(context.Context, *FeatureGottenEvent) error
	EmitListedEvent(context.Context, *FeatureListedEvent) error
}
