package auth

import (
	// "context"

	// "github.com/igdanit/sso/protos/gen/go/auth"
	sso "github.com/igdanit/sso/protos/gen/go/auth"
)

type AuthService struct {
	sso.UnimplementedAuthServer
}

// func (service AuthService) Register(ctx context.Context, res *auth.RegisterRequest) (auth.RegisterResponse, error) {

// }
