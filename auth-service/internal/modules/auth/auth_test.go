package auth

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"auth-service/internal/models"
	"auth-service/mocks"
)

func TestAuth_Login(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	userID := "test-user-id"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	user := models.User{
		ID:       userID,
		HashPass: hashedPassword,
	}

	t.Run("Success", func(t *testing.T) {
		mockUserGetter := new(mocks.UserGetter)
		mockRefreshTokenSaver := new(mocks.RefreshTokenSaver)
		mockUserGetter.On("GetUser", ctx, userID).Return(user, nil)
		mockRefreshTokenSaver.On("SaveRefreshToken", ctx, mock.Anything, user.ID).Return(nil)

		auth := New(logger, nil, mockUserGetter, mockRefreshTokenSaver)

		token, err := auth.Login(ctx, userID, password)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		mockUserGetter.AssertExpectations(t)
		mockRefreshTokenSaver.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserGetter := new(mocks.UserGetter)
		mockUserGetter.On("GetUser", ctx, userID).Return(models.User{}, errors.New("user not found"))

		auth := New(logger, nil, mockUserGetter, nil)

		token, err := auth.Login(ctx, userID, password)

		assert.Error(t, err)
		assert.Empty(t, token)
		mockUserGetter.AssertExpectations(t)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mockUserGetter := new(mocks.UserGetter)
		invalidPassword := "invalid_password"
		mockUserGetter.On("GetUser", ctx, userID).Return(user, nil)

		auth := New(logger, nil, mockUserGetter, nil)

		token, err := auth.Login(ctx, userID, invalidPassword)

		assert.Error(t, err)
		assert.Empty(t, token)
		mockUserGetter.AssertExpectations(t)
	})
}

func TestAuth_Register(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	email := "test@example.com"
	password := "password123"
	userType := "user"

	t.Run("Success", func(t *testing.T) {
		mockUserRegister := new(mocks.UserRegister)
		mockUserRegister.On("SaveUser", ctx, email, mock.Anything, userType).Return("test-uuid", nil)

		auth := New(logger, mockUserRegister, nil, nil)

		uuid, err := auth.Register(ctx, email, password, userType)

		assert.NoError(t, err)
		assert.Equal(t, "test-uuid", uuid)
		mockUserRegister.AssertExpectations(t)
	})

	t.Run("User Exists", func(t *testing.T) {
		mockUserRegister := new(mocks.UserRegister)
		mockUserRegister.On("SaveUser", ctx, email, mock.Anything, userType).Return("", errors.New("user already exists"))

		auth := New(logger, mockUserRegister, nil, nil)

		uuid, err := auth.Register(ctx, email, password, userType)

		assert.Error(t, err)
		assert.Empty(t, uuid)
		mockUserRegister.AssertExpectations(t)
	})
}

func TestAuth_DummyLogin(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	userType := "dummy-user"

	t.Run("Success", func(t *testing.T) {
		auth := New(logger, nil, nil, nil)

		token, err := auth.DummyLogin(ctx, userType)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}
