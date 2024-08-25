package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	str "sso/internal/storage"
)

var (
	isAdminQuery = `select is_admin from users where id = ?`
)

func (storage *Storage) IsAdmin(ctx context.Context, uid int64) (bool, error) {
	const op = "storage.sqlite.IsAdmin"

	prepare, err := storage.db.Prepare(isAdminQuery)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := prepare.QueryRowContext(ctx, uid)

	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, str.ErrAppNotFound)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
