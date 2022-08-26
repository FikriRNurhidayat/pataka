package domain

import "context"

const (
	AUDIENCE_CREATED_EVENT = "audiences.created"
	AUDIENCE_UPDATED_EVENT = "audiences.updated"
	AUDIENCE_GOTTEN_EVENT  = "audiences.gotten"
	AUDIENCE_DELETED_EVENT = "audiences.deleted"
	AUDIENCE_LISTED_EVENT  = "audiences.listed"
)

type AudienceCreatedEvent Event[Audience]
type AudienceUpdatedEvent Event[Audience]
type AudienceGottenEvent Event[Audience]
type AudienceListedEvent Event[[]Audience]
type AudienceDeletedEvent Event[Audience]

type AudienceEventEmitter interface {
	EmitCreatedEvent(context.Context, *AudienceCreatedEvent) error
	EmitUpdatedEvent(context.Context, *AudienceUpdatedEvent) error
	EmitDeletedEvent(context.Context, *AudienceDeletedEvent) error
	EmitGottenEvent(context.Context, *AudienceGottenEvent) error
	EmitListedEvent(context.Context, *AudienceListedEvent) error
}
