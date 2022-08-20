package auth

import "context"

type Authenticatable interface {
	Valid(context.Context) error
}
