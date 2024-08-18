package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/models"
)


func NewToken(user models.User,  signingKey []byte, duration time.Duration) (string, error) {
	op := "jwt.NewToken"
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["user_type"] = user.Type
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	return tokenString, nil
}

func NewRefreshToken(lastChar string, signingKey []byte, ttl time.Duration) (refreshToken string, err error) {
	op := "jwt.NewRefreshToken"
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["last_char"] = lastChar

	refreshToken, err = token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	return refreshToken, nil
}

func ValidateToken(accessToken string, signingKey []byte) (bool, string, string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return false, "", "", err
	}

	// Проверка и извлечение UID из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uidString, ok := claims["uid"].(string)
		if !ok {
			return false, "", "", fmt.Errorf("invalid uid claim in token")
		}
		userType, ok := claims["user_type"].(string)
		if !ok {
			return false, "", "", fmt.Errorf("invalid usertype claim in token")
		}
		return true, uidString, userType, nil
	}

	return false, "", "", fmt.Errorf("invalid token")
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

func IdFromJWT(accessToken string) (uid string, err error) {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		uid, ok := claims["uid"] 
		if !ok {
			return "", fmt.Errorf("uid must be a string, got %T", claims["uid"]) // Создаем новую ошибку, которая объясняет проблему
		}
		strUid := fmt.Sprint(uid)
		return strUid, nil
	}
	return "", fmt.Errorf("invalid JWT claims, unable to assert to MapClaims")
}
