package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"
	"strconv"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	grpcServer := grpc.NewServer()

	authgrpc.Register(grpcServer)

	return &App{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}

func (app *App) MustStart() {
	if err := app.Start(); err != nil {
		panic(err)
	}
}

func (app *App) Start() error {
	const op = "grpcapp.Start"

	log := app.log.With(
		slog.String("op", op),
		slog.String("port", strconv.Itoa(app.port)),
	)

	log.Info("starting grpc server")

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server started", slog.String("address", lis.Addr().String()))

	if err := app.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"

	app.log.With(slog.String("op", op)).Info("stopping grpc server")

	app.grpcServer.GracefulStop()
}
