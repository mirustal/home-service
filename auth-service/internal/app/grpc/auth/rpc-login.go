package grpcauth

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbErr "auth-service/internal/adapters/db"
	authgrpc "auth-service/pkg/pb"
)

func (s *serverAPI) Login(ctx context.Context, req *authgrpc.LoginRequest) (*authgrpc.LoginResponse, error) {
	userid, pass, err := validateLogin(req)
	if pass == "" {
		return nil, status.Error(codes.InvalidArgument, "pass is required")
	}

	token, err := s.auth.Login(ctx, userid, pass)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid userid or password")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to login: %w", err))
	}

	return &authgrpc.LoginResponse{
		Token: token,
	}, nil
}

func validateLogin(req *authgrpc.LoginRequest) (string, string, error) {
	userid := req.GetId()
	if userid == "" {
		return "", "", status.Error(codes.InvalidArgument, "userid is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return "", "", status.Error(codes.InvalidArgument, "password is required")
	}

	return userid, pass, nil
}
