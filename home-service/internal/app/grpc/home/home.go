package home

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	authclient "home-service/internal/client/auth-service"
	"home-service/internal/models"
	"home-service/pkg/pb"
)

type serverAPI struct {
	pb.UnimplementedHouseServiceServer
	home       Home
	authClient authclient.AuthClient
}

func Register(gRPC *grpc.Server, home Home, authCleint authclient.AuthClient) {
	pb.RegisterHouseServiceServer(gRPC, &serverAPI{
		home:       home,
		authClient: authCleint,
	})
}

const BearerPrefix = "Bearer "

type Home interface {
	CreateFlat(ctx context.Context, house_id int, price int, rooms int) (flat models.Flat, err error)
	CreateHouse(ctx context.Context, address string, year int, developer string) (house models.House, err error)
	GetFlatsInHouse(ctx context.Context, id int, userType string) (flat []models.Flat, err error)
	SubscribeToHouse(ctx context.Context, id int, email string) (err error)
	UpdateFlat(ctx context.Context, id int, status string) (models.Flat, error)
}


func (s *serverAPI) AuthCheck(ctx context.Context) (string, string, error) {
	reqToken, err := ExtractAccessToken(ctx)
	if err != nil {
		return "", "", status.Error(codes.Unauthenticated, "access token not found")
	}

	token, err := ValidateAccessToken(reqToken)
	if err != nil {
		return "", "", status.Error(codes.InvalidArgument, "invalid token")
	}


	isValid, uid, userStatus, err := s.authClient.ValidateToken(ctx, token)
	if err != nil {
		return "", "", fmt.Errorf("failed to validate token: %v", err)
	}

	if !isValid {
		return "", "", fmt.Errorf("invalid access token")
	}

	return uid, userStatus, nil
}

func ExtractAccessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	if tokens, exists := md["authorization"]; exists && len(tokens) > 0 {
		return tokens[0], nil
	}

	return "", fmt.Errorf("access token not found")
}

func ValidateAccessToken(token string) (string, error) {
	if token == "" {
		return "", status.Error(codes.InvalidArgument, "invalid token")
	}

	if !strings.HasPrefix(token, BearerPrefix) {
		return "", status.Error(codes.InvalidArgument, "invalid token")
	}

	return token, nil
}

func IsModerator(userType string) (isModerator bool) {
	if userType == "moderator" {
		isModerator = true
}
return isModerator
}