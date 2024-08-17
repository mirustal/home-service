package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/models"
)

const secret_key = "hello"

func NewToken(user models.User, duration time.Duration) (string, error) {
	op := "jwt.NewToken"
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["user_type"] = user.Type
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	return tokenString, nil
}

func NewRefreshToken(lastChar string, app models.App, duration time.Duration) (refreshToken string, err error) {
	op := "jwt.NewRefreshToken"
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["last_char"] = lastChar

	refreshToken, err = token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	return refreshToken, nil
}

func ValidateToken(accessToken string, signingKey []byte) (bool, int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return false, 0, err
	}

	// Проверка и извлечение UID из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uidFloat, ok := claims["uid"].(float64)
		if !ok {
			return false, 0, fmt.Errorf("invalid uid claim in token")
		}
		return true, int64(uidFloat), nil
	}

	return false, 0, fmt.Errorf("invalid token")
}

func CheckRefreshToken(refreshToken, accessToken string, signingKey []byte) error {
	if len(accessToken) < 6 {
		return fmt.Errorf("invalid access token length")
	}

	lastChar := accessToken[len(accessToken)-6:]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		forCheck, ok := claims["last_char"].(string)
		if !ok {
			return fmt.Errorf("invalid last_char claim in refresh token")
		}
		if lastChar != forCheck {
			return fmt.Errorf("token pair mismatch")
		}
		return nil
	}

	return fmt.Errorf("invalid refresh token")
}
