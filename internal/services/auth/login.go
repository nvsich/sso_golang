package auth

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"sso/internal/lib/jwt"
	"sso/internal/storage"
)

func (auth *Auth) Login(ctx context.Context, email, password string, appID int64) (string, error) {
	const op = "auth.Login"

	log := auth.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attempting to login")

	user, err := auth.userRepo.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			auth.log.Info("user not found", err)
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		auth.log.Error("failed to get user", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		auth.log.Info("invalid credentials", err)
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := auth.appRepo.App(ctx, appID)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			auth.log.Info("app not found", err)
			return "", fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}

		auth.log.Error("failed to get app", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully logged in")

	token, err := jwt.NewToken(user, app, auth.tokenTTL)
	if err != nil {
		auth.log.Error("failed to create token", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}
