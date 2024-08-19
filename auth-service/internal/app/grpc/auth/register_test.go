package grpcauth

import (
	"context"
	"fmt"
	"testing"

	"auth-service/mocks"
	authgrpc "auth-service/pkg/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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
