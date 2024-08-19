package jwt

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const BearerPrefix = "Bearer "

func IdFromJWT(accessToken string) (uid string, err error) {
	accessToken, ok := strings.CutPrefix(accessToken, BearerPrefix)
		if !ok {
			return "", fmt.Errorf("token not valid") 
		}
		
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		uid, ok := claims["uid"]
		if !ok {
			return "", fmt.Errorf("uid must be a string, got %T", claims["uid"])
		}
		strUid := fmt.Sprint(uid)
		return strUid, nil
	}
	return "", fmt.Errorf("invalid JWT claims, unable to assert to MapClaims")
}
