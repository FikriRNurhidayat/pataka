package authentication

import (
	"context"
	"strings"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/inspector"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authentication struct {
	SecretKey string
}

func (s *Authentication) CreateToken(scopes ...string) (string, error) {
	claims := &domain.Claims{
		StandardClaims: jwt.StandardClaims{},
		Scopes:         scopes,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Authentication) Call(ctx context.Context, scopes ...string) (*domain.Claims, error) {
	tokenString, err := s.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := s.GetClaims(tokenString)
	if err != nil {
		return nil, err
	}

	if !inspector.IsEmptySlice(scopes) {
		allowed := s.IsScopeAllowed(claims, scopes...)
		if !allowed {
			return nil, status.Error(codes.PermissionDenied, "Access not allowed.")
		}
	}

	return claims, nil
}

func (s *Authentication) GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "No API Key.")
	}

	headers := md.Get("authorization")

	if len(headers) == 0 {
		return "", status.Error(codes.Unauthenticated, "No API Key in header.")
	}

	bearerToken := headers[0]

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", status.Error(codes.Unauthenticated, "Authorization must be in Bearer format.")
	}

	return strings.ReplaceAll(bearerToken, "Bearer ", ""), nil
}

func (s *Authentication) GetClaims(tokenString string) (*domain.Claims, error) {
	claims := &domain.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token.")
	}

	return claims, nil
}

func (s *Authentication) IsScopeAllowed(claims *domain.Claims, scopes ...string) bool {
	for _, cs := range claims.Scopes {
		for _, s := range scopes {
			if s == cs {
				return true
			}
		}
	}

	return false
}

func New(secretKey string) domain.Authenticatable {
	return &Authentication{
		SecretKey: secretKey,
	}
}
