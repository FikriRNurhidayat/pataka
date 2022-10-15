package domain

import (
	"context"

	"github.com/golang-jwt/jwt"
)

type Authenticatable interface {
	Call(ctx context.Context, scopes ...string) (*Claims, error)
	CreateToken(scopes ...string) (string, error)
	GetClaims(tokenString string) (*Claims, error)
	GetToken(ctx context.Context) (string, error)
	IsScopeAllowed(claims *Claims, scopes ...string) bool
}

type Claims struct {
	jwt.StandardClaims
	Scopes []string `json:"scopes"`
}

var (
	ReadAudienceScope  string = "read:audience"
	WriteAudienceScope string = "write:audience"
	ReadFeatureScope   string = "read:feature"
	WriteFeatureScope  string = "write:feature"
)

var AllScope []string = []string{ReadAudienceScope, WriteAudienceScope, ReadFeatureScope, WriteFeatureScope}
