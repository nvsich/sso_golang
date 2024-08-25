package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sso/internal/domain/models"
	str "sso/internal/storage"
)

var (
	getUserQuery = `select id, email, pass_hash from users where email = ?`
)

func (storage *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.sqlite.User"

	prepare, err := storage.db.Prepare(getUserQuery)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := prepare.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PassHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, str.ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
