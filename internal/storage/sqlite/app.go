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
	getAppQuery = `select id, name, secret from apps where id = ?`
)

func (storage *Storage) App(ctx context.Context, appID int64) (models.App, error) {
	const op = "storage.sqlite.App"

	prepare, err := storage.db.Prepare(getAppQuery)
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	row := prepare.QueryRowContext(ctx, appID)

	var app models.App
	err = row.Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, str.ErrAppNotFound)
		}

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
