package grpcauth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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
