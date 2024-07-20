package auth

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"sso/internal/storage"
)

func (auth *Auth) Register(ctx context.Context, email, password string) (int64, error) {
	const op = "auth.Register"

	log := auth.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error("failed to hash password", err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	userID, err := auth.userRepo.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Info("user exists", err)
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		log.Error("failed to save user", err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully registered user")

	return userID, nil
}
