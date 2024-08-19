package grpcauth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbErr "auth-service/internal/adapters/db"
	authgrpc "auth-service/pkg/pb"
)

func (s *serverAPI) Register(ctx context.Context, req *authgrpc.RegisterRequest) (*authgrpc.RegisterResponse, error) {
	email, pass, userType, err := validateRegister(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "validate error")
	}

	userid, err := s.auth.Register(ctx, email, pass, userType)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "The user already exists")
		}
		return nil, status.Error(codes.AlreadyExists, "The user already exists")
	}

	return &authgrpc.RegisterResponse{
		UserId: userid,
	}, nil
}

func validateRegister(req *authgrpc.RegisterRequest) (string, string, string, error) {
	email := req.GetEmail()
	if email == "" {
		return "", "", "", status.Error(codes.InvalidArgument, "email is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return "", "", "", status.Error(codes.InvalidArgument, "password is required")
	}

	userType := req.GetUserType()
	if userType != authgrpc.UserType_client && userType != authgrpc.UserType_moderator {
		return "", "", "", status.Error(codes.InvalidArgument, "invalid user type")
	}

	return email, pass, userType.String(), nil
}
