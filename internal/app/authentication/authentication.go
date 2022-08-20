package authentication

import "context"

type Authenticatable interface {
	Valid(context.Context) error
	GetToken(context.Context) (string, error)
	CreateToken() (string, error)
}
