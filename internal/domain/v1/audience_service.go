package domain

import "context"

type AudienceBulkCreatable interface {
	Call(context.Context, *BulkCreateAudiencesAudienceParams) (*BulkCreateAudiencesAudienceResult, error)
}

type AudienceCreatable interface {
	Call(context.Context, *CreateAudienceParams) (*CreateAudienceResult, error)
}

type AudienceGetable interface {
	Call(context.Context, *GetAudienceParams) (*GetAudienceResult, error)
}

type AudienceListable interface {
	Call(context.Context, *ListAudiencesParams) (*ListAudiencesResult, error)
}

type AudienceUpdateable interface {
	Call(context.Context, *UpdateAudienceParams) (*UpdateAudienceResult, error)
}

type AudienceDeletable interface {
	Call(context.Context, *DeleteAudienceParams) error
}
