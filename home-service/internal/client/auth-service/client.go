package authclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"home-service/pkg/pb"
)

// AuthClient интерфейс для клиента авторизации
type AuthClient interface {
	ValidateToken(ctx context.Context, token string) (bool, string, string, error)
}

// authClient реализация интерфейса AuthClient
type authClient struct {
	client pb.AuthClient
}

// NewAuthClient создает новый клиент авторизации
func NewAuthClient(address string) (AuthClient, error) {
	fmt.Println(address)
	op := "internal.client.Auth.NewAuthClient"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
	client := pb.NewAuthClient(conn)
	return &authClient{client: client}, nil
}

// ValidateToken выполняет валидацию токена через сервис авторизации
func (c *authClient) ValidateToken(ctx context.Context, token string) (bool, string, string, error) {
	op := "internal.client.Auth.ValidateToken"
	validateReq := &pb.ValidateRequest{AccessToken: token}
	validateResp, err := c.client.ValidateSession(ctx, validateReq)
	if err != nil {
		return false, "", "", fmt.Errorf("%s, %w", op, err)
	}
	return validateResp.GetIsValid(), validateResp.GetUid(), validateResp.GetUserType(), nil
}
