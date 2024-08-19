package grpcauth

import (
	"context"
	"testing"

	dbErr "auth-service/internal/adapters/db"
	"auth-service/mocks"
	authgrpc "auth-service/pkg/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRefreshSession(t *testing.T) {
	mockAuth := new(mocks.Auth)
	server := &serverAPI{auth: mockAuth}

	t.Run("success", func(t *testing.T) {
		mockAuth.On("RefreshSession", mock.Anything, "valid_access_token", "valid_refresh_token").
			Return("new_access_token", "new_refresh_token", nil)

		req := &authgrpc.RefreshRequest{AccessToken: "valid_access_token", RefreshToken: "valid_refresh_token"}
		resp, err := server.RefreshSession(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, "new_access_token", resp.NewAccessToken)
		assert.Equal(t, "new_refresh_token", resp.NewRefreshToken)
		mockAuth.AssertExpectations(t)
	})

	t.Run("invalid refresh token", func(t *testing.T) {
		mockAuth.On("RefreshSession", mock.Anything, "valid_access_token", "invalid_refresh_token").
			Return("", "", dbErr.ErrInvalidRefreshToken)

		req := &authgrpc.RefreshRequest{AccessToken: "valid_access_token", RefreshToken: "invalid_refresh_token"}
		resp, err := server.RefreshSession(context.Background(), req)

		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "invalid refresh token")
		mockAuth.AssertExpectations(t)
	})
}
