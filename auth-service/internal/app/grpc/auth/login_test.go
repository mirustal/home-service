package grpcauth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbErr "auth-service/internal/adapters/db"
	"auth-service/mocks"
	authgrpc "auth-service/pkg/pb"
)

func TestLogin(t *testing.T) {
	mockAuth := new(mocks.Auth)
	server := &serverAPI{auth: mockAuth}

	t.Run("success", func(t *testing.T) {
		mockAuth.On("Login", mock.Anything, "admin", "admin").Return("valid_token", nil)

		req := &authgrpc.LoginRequest{Id: "admin", Password: "admin"}
		resp, err := server.Login(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, "valid_token", resp.Token)
		mockAuth.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockAuth.On("Login", mock.Anything, "unknown_user", "valid_pass").Return("", dbErr.ErrUserNotFound)

		req := &authgrpc.LoginRequest{Id: "unknown_user", Password: "valid_pass"}
		resp, err := server.Login(context.Background(), req)

		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "invalid userid or password")
		mockAuth.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		mockAuth.On("Login", mock.Anything, "valid_user", "valid_pass").Return("", errors.New("some error"))

		req := &authgrpc.LoginRequest{Id: "valid_user", Password: "valid_pass"}
		resp, err := server.Login(context.Background(), req)

		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "failed to login")
		mockAuth.AssertExpectations(t)
	})
}
