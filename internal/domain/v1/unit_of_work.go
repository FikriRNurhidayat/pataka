package domain

import "context"

type Block func(Repository) error

type UnitOfWork interface {
	Do(context.Context, Block) error
}
