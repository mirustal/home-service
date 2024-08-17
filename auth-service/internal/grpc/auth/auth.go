package authgrpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbErr "auth-service/internal/adapters/db"
	authgrpc "auth-service/pkg/pb"
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

type Auth interface {
	Login(ctx context.Context, userid string, pass string) (token string, err error)
	Register(ctx context.Context, email string, pass string, userType string) (userid string, err error)
	DummyLogin(ctx context.Context, userType string) (token string, err error)
}

func (s *serverAPI) Login(ctx context.Context, req *authgrpc.LoginRequest) (*authgrpc.LoginResponse, error) {
	userid := req.UserId // ВЫНЕСИ ВАЛИДАЦИЮ В ОТДЕЛЬНЫЙ МТЕОД
	if userid == "" {
		return nil, status.Error(codes.InvalidArgument, "userid is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return nil, status.Error(codes.InvalidArgument, "pass is required")
	}

	token, err := s.auth.Login(ctx, userid, pass)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid userid or password")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to login: %v", err))
	}

	return &authgrpc.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *authgrpc.RegisterRequest) (*authgrpc.RegisterResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return nil, status.Error(codes.InvalidArgument, "pass is required")
	}

	userType := req.UserType

	if userType != *authgrpc.UserType_client.Enum() {
		return nil, status.Error(codes.InvalidArgument, "userType is required")
	}

	userid, err := s.auth.Register(ctx, email, pass, userType.String())
	if err != nil {
		if errors.Is(err, dbErr.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "The user already exists")
		}
		return nil, status.Error(codes.Internal, "failed to register")
	}

	return &authgrpc.RegisterResponse{
		UserId: userid,
	}, nil
}

func (s *serverAPI) DummyLogin(ctx context.Context, req *authgrpc.DummyLoginRequest) (*authgrpc.DummyLoginResponse, error) {
	token, err := s.auth.DummyLogin(ctx, req.GetUserType().String())
	if err != nil{ 
		return nil, status.Error(codes.Internal,fmt.Sprintf("failed to login: %v", err))
	} 

	return &authgrpc.DummyLoginResponse{
		Token: token,
	}, nil
}
