package sqlite

import (
	"context"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	str "sso/internal/storage"
)

var (
	saveUserQuery = `insert into users(email, pass_hash) values (?, ?)`
)

func (storage *Storage) SaveUser(ctx context.Context, email, passHash []byte) (int64, error) {
	const op = "storage.sqlite.SaveUser"
	prepare, err := storage.db.Prepare(saveUserQuery)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := prepare.ExecContext(ctx, email, passHash)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, str.ErrUserExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
