package grpcauth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbErr "auth-service/internal/adapters/db"
	authgrpc "auth-service/pkg/pb"
)

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
