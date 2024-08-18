package grpcauth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	dbErr "auth-service/internal/adapters/db"
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

type Auth interface {
	Login(ctx context.Context, userid string, pass string) (token string, err error)
	Register(ctx context.Context, email string, pass string, userType string) (userid string, err error)
	DummyLogin(ctx context.Context, userType string) (token string, err error)
	RefreshSession(ctx context.Context, refreshToken string, accessToken string) (string, string, error)
	ValidateSession(ctx context.Context, accessToken string) (isValid bool, uid string, userType string, err error)
}

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

		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to login: %v", err))
	}

	return &authgrpc.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *authgrpc.RegisterRequest) (*authgrpc.RegisterResponse, error) {
	email, pass, userType, err := validateRegister(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%v", err))
	}

	userid, err := s.auth.Register(ctx, email, pass, userType)
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
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to login: %v", err))
	}

	return &authgrpc.DummyLoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) ValidateSession(ctx context.Context, req *authgrpc.ValidateRequest) (*authgrpc.ValidateResponse, error) {
	// Извлекаем токен из метаданных или тела запроса
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
		IsValid: isValid,
		Uid:     uid,
		UserType: userType,
	}, nil
}

func (s *serverAPI) RefreshSession(ctx context.Context, req *authgrpc.RefreshRequest) (*authgrpc.RefreshResponse, error) {
	newAccessToken, newRefreshToken, err := s.auth.RefreshSession(ctx, req.AccessToken, req.GetRefreshToken())
	if err != nil {
		if errors.Is(err, dbErr.ErrInvalidRefreshToken) {
			return nil, status.Error(codes.InvalidArgument, "invalid refresh token")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authgrpc.RefreshResponse{
		NewAccessToken:  newAccessToken,
		NewRefreshToken: newRefreshToken,
	}, nil
}

func ExtractAccessToken(ctx context.Context, reqTokenFromBody string) (string, error) {
	// Попытка извлечь токен из метаданных контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if tokens, exists := md["authorization"]; exists && len(tokens) > 0 {
			return tokens[0], nil
		}
	}

	// Если токен не найден в метаданных, проверить тело запроса
	if reqTokenFromBody != "" {
		return reqTokenFromBody, nil
	}

	return "", fmt.Errorf("access token not found")
}



func validateLogin(req *authgrpc.LoginRequest) (string, string, error) {
	userid := req.GetId()
	if userid == "" {
		return "", "", status.Error(codes.InvalidArgument, "userid is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return "", "", status.Error(codes.InvalidArgument, "pass is required")
	}

	return userid, pass, nil
}

func validateRegister(req *authgrpc.RegisterRequest) (string, string, string, error) {
	email := req.GetEmail()
	if email == "" {
		return  "", "", "", status.Error(codes.InvalidArgument, "email is required")
	}

	pass := req.GetPassword()
	if pass == "" {
		return "", "", "", status.Error(codes.InvalidArgument, "pass is required")
	}

	userType := req.GetUserType()
	return email, pass, userType.String(), nil
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
