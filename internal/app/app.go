package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: initialize storage

	// TODO: init auth service

	grpcServer := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcServer,
	}
}
