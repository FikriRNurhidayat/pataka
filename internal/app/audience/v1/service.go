package audience

import "context"

type BulkCreateable interface {
	Call(context.Context, *BulkCreateParams) (*BulkCreateResult, error)
}

type Createable interface {
	Call(context.Context, *CreateParams) (*CreateResult, error)
}

type Getable interface {
	Call(context.Context, *GetParams) (*GetResult, error)
}

type Listable interface {
	Call(context.Context, *ListParams) (*ListResult, error)
}

type Updatable interface {
	Call(context.Context, *UpdateParams) (*UpdateResult, error)
}

type Deletable interface {
	Call(context.Context, *DeleteParams) error
}
