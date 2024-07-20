package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sso/internal/storage"
)

func (auth *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "auth.IsAdmin"

	log := auth.log.With(
		slog.String("op", op),
		slog.String("userID", fmt.Sprint(userID)),
	)

	log.Info("checking if user is admin")

	isAdmin, err := auth.adminRepo.IsAdmin(ctx, userID)

	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			auth.log.Info("user not found", err)
			return false, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}

		auth.log.Error("failed to check if user is admin", err)
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully checked if user is admin", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}
