package grpcauth

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authgrpc "auth-service/pkg/pb"
)

func (s *serverAPI) ValidateSession(ctx context.Context, req *authgrpc.ValidateRequest) (*authgrpc.ValidateResponse, error) {
	reqToken, err := ExtractAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "access token not found")
	}

	token, err := validateAccessToken(reqToken)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token")
	}

	isValid, uid, userType, err := s.auth.ValidateSession(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	return &authgrpc.ValidateResponse{
		IsValid:  isValid,
		Uid:      uid,
		UserType: userType,
	}, nil
}

func validateAccessToken(token string) (string, error) {
	if token == "" {
		return "", status.Error(codes.InvalidArgument, "invalid token")
	}

	if !strings.HasPrefix(token, bearerPrefix) {
		return "", status.Error(codes.InvalidArgument, "invalid token")
	}

	return strings.TrimPrefix(token, bearerPrefix), nil
}
