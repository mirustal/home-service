package grpcauth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	authgrpc "auth-service/pkg/pb"
)

const (
	bearerPrefix = "Bearer "
)

type serverAPI struct {
	authgrpc.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	authgrpc.RegisterAuthServer(gRPC, &serverAPI{
		auth: auth,
	})
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=Auth --with-expecter=true
type Auth interface {
	Login(ctx context.Context, userid string, pass string) (token string, err error)
	Register(ctx context.Context, email string, pass string, userType string) (userid string, err error)
	DummyLogin(ctx context.Context, userType string) (token string, err error)
	RefreshSession(ctx context.Context, refreshToken string, accessToken string) (string, string, error)
	ValidateSession(ctx context.Context, accessToken string) (isValid bool, uid string, userType string, err error)
}

func ExtractAccessToken(ctx context.Context, reqTokenFromBody string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if tokens, exists := md["authorization"]; exists && len(tokens) > 0 {
			return tokens[0], nil
		}
	}

	if reqTokenFromBody != "" {
		return reqTokenFromBody, nil
	}

	return "", fmt.Errorf("access token not found")
}
