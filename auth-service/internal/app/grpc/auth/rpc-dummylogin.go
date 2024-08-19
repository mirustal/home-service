package grpcauth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authgrpc "auth-service/pkg/pb"
)

func (s *serverAPI) DummyLogin(ctx context.Context, req *authgrpc.DummyLoginRequest) (*authgrpc.DummyLoginResponse, error) {
	userType := req.GetUserType()

	if userType == authgrpc.UserType_UNKNOWN {
		return nil, status.Error(codes.InvalidArgument, "user type not be empty")
	}
	
	if userType != authgrpc.UserType_client && userType != authgrpc.UserType_moderator {
		return nil, status.Error(codes.InvalidArgument, "invalid user type")
	}

	token, err := s.auth.DummyLogin(ctx, userType.String())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to login: %v", err))
	}

	return &authgrpc.DummyLoginResponse{
		Token: token,
	}, nil
}
