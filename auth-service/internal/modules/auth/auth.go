package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	dbErr "auth-service/internal/adapters/db"
	"auth-service/internal/models"
	"auth-service/pkg/jwt"
)

type Auth struct {
	log                 *slog.Logger
	userSaver           UserRegister
	userProvider        UserGetter
	refreshTokenSaver   RefreshTokenSaver
	refreshTokenChecker RefreshTokenChecker
}

func New(log *slog.Logger, userSave UserRegister, userProvider UserGetter, refreshTokenSavver RefreshTokenSaver) *Auth {
	return &Auth{
		log:               log,
		userSaver:         userSave,
		userProvider:      userProvider,
		refreshTokenSaver: refreshTokenSavver,
	}
}

var secret_key = os.Getenv("SECRET_KEY")


//go:generate go run github.com/vektra/mockery/v2@latest --name=UserRegister --with-expecter=true
type UserRegister interface {
	SaveUser(ctx context.Context, email string, hashPass []byte, userType string) (uuid string, err error)
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=RefreshTokenSaver --with-expecter=true
type RefreshTokenSaver interface {
	SaveRefreshToken(ctx context.Context, refreshToken string, uid string) error
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=RefreshTokenChecker --with-expecter=true

type RefreshTokenChecker interface {
	CheckRefreshToken(refreshToken string) error
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=UserGetter --with-expecter=true
type UserGetter interface {
	GetUser(ctx context.Context, userID string) (models.User, error)
}

func (a *Auth) Login(ctx context.Context, userid string, password string) (string, error) {
	const op = "internal.services.auth.Login"
	log := a.log.With(
		slog.String("op", op),
	)
	user, err := a.userProvider.GetUser(ctx, userid)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserNotFound) {
			log.Warn("user not found")
			return "", fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to get user")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.HashPass, []byte(password)); err != nil {
		log.Warn("password incorrect")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err := jwt.NewToken(user, []byte(secret_key), time.Hour)
	if err != nil {
		log.Error("failed to generated token")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println(accessToken)
	lastChar := accessToken[len(accessToken)-6:]

	refreshToken, err := jwt.NewRefreshToken(lastChar, []byte(secret_key), time.Hour*24)
	if err != nil {
		a.log.Error("failed to generate refresh token", err)

		return "", fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println(refreshToken)

	if err := a.refreshTokenSaver.SaveRefreshToken(context.Background(), refreshToken, user.ID); err != nil {
		a.log.Error("failed to save refresh token", err)
	}

	log.Info("user logged in successful")
	return accessToken, nil
}


func (a *Auth) DummyLogin(ctx context.Context, userType string) (string, error) {
	const op = "internal.services.auth.dummyLogin"
	log := a.log.With(
		slog.String("op", op),
	)
	user := models.User{
		Type: userType,
	}
	token, err := jwt.NewToken(user, []byte(secret_key), time.Hour)
	if err != nil {
		log.Error("failed to generated token")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successful")
	return token, nil
}

func (a *Auth) Register(ctx context.Context, email string, pass string, userType string) (string, error) {
	const op = "internal.services.auth.RegisterNewUser"
	log := a.log.With(
		slog.String("op", op),
	)

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Error("failed to generate password ")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	uuid, err := a.userSaver.SaveUser(context.Background(), email, passHash, userType)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserExists) {
			log.Warn("user Exist")
			return "", fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to save user in database")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User register")
	return uuid, nil
}

func (a *Auth) RefreshSession(ctx context.Context, refreshToken string, accessToken string) (string, string, error) {
	const op = "auth.RefreshSession"

	log := a.log.With(slog.String("op", op))

	log.Info("refresh user token")

	err := jwt.CheckRefreshToken(refreshToken, accessToken, []byte(secret_key))
	if err != nil {
		log.Error("failed to validate token pair", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	uid, err := jwt.IdFromJWT(accessToken)
	if err != nil {
		log.Error("failed to extract uid from token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	err = a.refreshTokenChecker.CheckRefreshToken(refreshToken)
	if err != nil {
		a.log.Error("invalid refresh token", err)
		return "", "", err
	}

	var user models.User
	user.ID = string(uid)

	newAccessToken, err := jwt.NewToken(user, []byte(secret_key), time.Hour)
	if err != nil {
		a.log.Error("failed to generate access token", err)

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	lastChar := newAccessToken[len(newAccessToken)-6:]

	newRefreshToken, err := jwt.NewRefreshToken(lastChar, []byte(secret_key), time.Hour)
	if err != nil {
		a.log.Error("failed to generate refresh token", err)

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if err := a.refreshTokenSaver.SaveRefreshToken(context.Background(), newRefreshToken, user.ID); err != nil {
		a.log.Error("failed to save refresh token", err)
	}

	log.Info("user token successfully refreshed")

	return newAccessToken, newRefreshToken, nil
}

func (a *Auth) ValidateSession(ctx context.Context, accessToken string) (isValid bool, uid string, userType string, err error) {
	const op = "auth.ValidateSession"

	log := a.log.With(slog.String("op", op))

	log.Info("validate user token")

	isValid, uid, userType, err = jwt.ValidateToken(accessToken, []byte(secret_key))
	if err != nil {
		log.Error("failed validate token", err)
		return false, "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("token successfully validate")

	return isValid, fmt.Sprint(uid), userType, nil
}
