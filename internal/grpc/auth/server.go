package auth

import (
	"context"
	ssov1 "github.com/nvsich/sso_protos/gen/go/sso"
	"google.golang.org/grpc"
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (string, error)

	Register(ctx context.Context, email string, password string) (int64, error)

	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer

	auth Auth
}

func Register(grpc *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(grpc, &serverAPI{auth: auth})
}

const (
	emptyValue = 0
)
