package authentication

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthenticationService struct {
	SecretKey string
}

func (s *AuthenticationService) CreateToken() (string, error) {
	return jwt.New(jwt.SigningMethodHS256).SignedString([]byte(s.SecretKey))
}

func (s *AuthenticationService) GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "No API Key")
	}

	headers := md.Get("authorization")

	if len(headers) == 0 {
		return "", status.Error(codes.Unauthenticated, "No API Key in header")
	}

	bearerToken := headers[0]

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", status.Error(codes.Unauthenticated, "Authorization must be in Bearer format.")
	}

	return strings.ReplaceAll(bearerToken, "Bearer ", ""), nil
}

func (s *AuthenticationService) Valid(ctx context.Context) error {
	tokenString, err := s.GetToken(ctx)
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})

	return nil
}

func NewAuthenticaticationService(secretKey string) Authenticatable {
	return &AuthenticationService{
		SecretKey: secretKey,
	}
}
