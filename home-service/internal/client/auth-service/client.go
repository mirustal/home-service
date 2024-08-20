package authclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"home-service/pkg/pb"
)


//go:generate go run github.com/vektra/mockery/v2@latest --name=Home --with-expecter=true
type AuthClient interface {
	ValidateToken(ctx context.Context, token string) (bool, string, string, error)
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=Home --with-expecter=true
type authClient struct {
	client pb.AuthClient
}

func NewAuthClient(address string) (AuthClient, error) {
	op := "internal.client.Auth.NewAuthClient"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
	client := pb.NewAuthClient(conn)
	return &authClient{client: client}, nil
}

func (c *authClient) ValidateToken(ctx context.Context, token string) (bool, string, string, error) {
	op := "internal.client.Auth.ValidateToken"
	validateReq := &pb.ValidateRequest{AccessToken: token}
	validateResp, err := c.client.ValidateSession(ctx, validateReq)
	if err != nil {
		return false, "", "", fmt.Errorf("%s, %w", op, err)
	}
	return validateResp.GetIsValid(), validateResp.GetUid(), validateResp.GetUserType(), nil
}
