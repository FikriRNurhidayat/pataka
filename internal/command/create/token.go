package create

import (
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
)

type CreateTokenCmd struct {
	authentication domain.Authenticatable
}

func (c *CreateTokenCmd) Call(scopes ...string) (string, error) {
	token, err := c.authentication.CreateToken(scopes...)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewCreateTokenCmd(authentication domain.Authenticatable) *CreateTokenCmd {
	return &CreateTokenCmd{
		authentication: authentication,
	}
}
