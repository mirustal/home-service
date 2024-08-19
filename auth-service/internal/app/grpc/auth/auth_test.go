package grpcauth

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbErr "auth-service/internal/adapters/db"
	"auth-service/mocks"
	authgrpc "auth-service/pkg/pb"
)

func TestDummyLogin(t *testing.T) {
	mockAuth := new(mocks.Auth)
	server := &serverAPI{auth: mockAuth}

	t.Run("success", func(t *testing.T) {
		mockAuth.On("DummyLogin", mock.Anything, "client").Return("mocked_token", nil)

		req := &authgrpc.DummyLoginRequest{UserType: authgrpc.UserType_client}
		resp, err := server.DummyLogin(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, "mocked_token", resp.Token)
		mockAuth.AssertExpectations(t)
	})

}

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

	t.Run("invalid password", func(t *testing.T) {
		req := &authgrpc.LoginRequest{Id: "valid_user"}
		resp, err := server.Login(context.Background(), req)

		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "pass is required")
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

func TestRegister(t *testing.T) {
	mockAuth := new(mocks.Auth)
	server := &serverAPI{auth: mockAuth}

	t.Run("success", func(t *testing.T) {
		mockAuth.On("Register", mock.Anything, "user@example.com", "valid_pass", "client").
			Return("new_user_id", nil)

		req := &authgrpc.RegisterRequest{Email: "user@example.com", Password: "valid_pass", UserType: authgrpc.UserType_client}
		resp, err := server.Register(context.Background(), req)

		require.NoError(t, err)
		fmt.Println(resp.UserId)
		assert.Equal(t, "new_user_id", resp.UserId)
		mockAuth.AssertExpectations(t)
	})
}
