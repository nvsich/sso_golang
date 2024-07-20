package auth

import (
	"context"
	"errors"
	"log/slog"
	"sso/internal/domain/models"
	"time"
)

type Auth struct {
	log       *slog.Logger
	userRepo  UserRepo
	adminRepo AdminRepo
	appRepo   AppRepo
	tokenTTL  time.Duration
}

type UserRepo interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (int64, error)

	User(ctx context.Context, email string) (models.User, error)
}

type AdminRepo interface {
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppRepo interface {
	App(ctx context.Context, appId int64) (models.App, error)
}

func New(
	log *slog.Logger,
	userRepo UserRepo,
	adminRepo AdminRepo,
	appRepo AppRepo,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:       log,
		userRepo:  userRepo,
		adminRepo: adminRepo,
		appRepo:   appRepo,
		tokenTTL:  tokenTTL,
	}
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAppNotFound        = errors.New("app not found")
	ErrUserExists         = errors.New("user already exists")
)
