package suite

import (
	"context"
	ssov1 "github.com/nvsich/sso_protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"sso/internal/config"
	"strconv"
	"testing"
)

const (
	testConfigPath = "../config/local.yaml"
	grpcHost       = "0.0.0.0"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath(testConfigPath)

	ctx, cancelFunc := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelFunc()
	})

	conn, err := grpc.DialContext(context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials())) // Используем insecure-коннект для тестов

	if err != nil {
		t.Fatalf("grpc connection failed: %v", err)
	}

	return ctx,
		&Suite{
			T:          t,
			Cfg:        cfg,
			AuthClient: ssov1.NewAuthClient(conn),
		}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
