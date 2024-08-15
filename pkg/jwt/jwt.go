package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"home-service/internal/models"
)

const secret_key = "hello" 

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"]= time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
