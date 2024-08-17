package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
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

func New(log *slog.Logger, userSave UserRegister, userProvider UserGetter) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSave,
		userProvider: userProvider,
	}
}

type UserRegister interface {
	SaveUser(ctx context.Context, email string, hashPass []byte, userType string) (uuid string, err error)
}

type RefreshTokenSaver interface {
	SaveRefreshToken(refreshToken string, uid int64) error
}

type RefreshTokenChecker interface {
	CheckRefreshToken(refreshToken string) error
}

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
			return "", fmt.Errorf("%s: %v", op, err)
		}

		log.Error("failed to get user")
		return "", fmt.Errorf("%s: %v", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.HashPass, []byte(password)); err != nil {
		log.Warn("password incorrect")
		return "", fmt.Errorf("%s: %v", op, err)
	}

	token, err := jwt.NewToken(user, time.Hour)
	if err != nil {
		log.Error("failed to generated token")
		return "", fmt.Errorf("%s: %v", op, err)
	}

	log.Info("user logged in successful")
	return token, nil
}

func (a *Auth) DummyLogin(ctx context.Context, userType string) (string, error) {
	const op = "internal.services.auth.dummyLogin"
	log := a.log.With(
		slog.String("op", op),
	)
	user := models.User{
		Type: userType,
	}
	token, err := jwt.NewToken(user, time.Hour)
	if err != nil {
		log.Error("failed to generated token")
		return "", fmt.Errorf("%s: %v", op, err)
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
		return "", fmt.Errorf("%s: %v", op, err)
	}

	uuid, err := a.userSaver.SaveUser(context.Background(), email, passHash, userType)
	if err != nil {
		if errors.Is(err, dbErr.ErrUserExists) {
			log.Warn("user Exist")
			return "", fmt.Errorf("%s: %v", op, err)
		}

		log.Error("failed to save user in database")
		return "", fmt.Errorf("%s: %v", op, err)
	}

	log.Info("User register")
	return uuid, nil
}
