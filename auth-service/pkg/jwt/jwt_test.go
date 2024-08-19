package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"

	"auth-service/internal/models"
)

func TestNewToken(t *testing.T) {

	signingKey := []byte("secret")
	user := models.User{
		ID:    "admin",
		Email: "user@example.com",
		Type:  "moderator",
	}
	duration := time.Hour

	tokenString, err := NewToken(user, signingKey, duration)
	assert.NoError(t, err)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return signingKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, user.ID, claims["uid"])
	assert.Equal(t, user.Email, claims["email"])
	assert.Equal(t, user.Type, claims["user_type"])
}

func TestNewRefreshToken(t *testing.T) {
	signingKey := []byte("secret")
	lastChar := "x"
	ttl := time.Hour

	refreshToken, err := NewRefreshToken(lastChar, signingKey, ttl)
	assert.NoError(t, err)

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return signingKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, lastChar, claims["last_char"])
}

func TestValidateToken(t *testing.T) {
	signingKey := []byte("secret")
	user := models.User{
		ID:    "12345",
		Email: "user@example.com",
		Type:  "admin",
	}
	duration := time.Hour

	tokenString, _ := NewToken(user, signingKey, duration)

	tests := []struct {
		name          string
		token         string
		expectedValid bool
		expectedUID   string
		expectedType  string
		expectedErr   error
	}{
		{
			name:          "Valid token",
			token:         tokenString,
			expectedValid: true,
			expectedUID:   user.ID,
			expectedType:  user.Type,
			expectedErr:   nil,
		},
		{
			name:          "Invalid token",
			token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIiwiZXhwIjoxNzI0MDI3MDUxLCJ1aWQiOiI4OGM1NGUxYS0xYTFiLTRkNzQtYWZhOS0wMGUxZjkzNDZiMDMiLCJ1c2VyX3R5cGUiOiJtb2RlcmF0b3IifQ.Zgk25SKD12HI2Rh46WnP1FzQlUfpqNHeKkOz7LSZhhs",
			expectedValid: false,
			expectedUID:   "",
			expectedType:  "",
			expectedErr:   jwt.ErrSignatureInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid, uid, userType, err := ValidateToken(test.token, signingKey)
			assert.Equal(t, test.expectedValid, valid)
			assert.Equal(t, test.expectedUID, uid)
			assert.Equal(t, test.expectedType, userType)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}
