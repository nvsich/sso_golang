package auth

import (
	"context"
	ssov1 "github.com/nvsich/sso_protos/gen/go/sso"
	"google.golang.org/grpc"
)

type Auth interface {
	Login(
		context context.Context,
		email string,
		password string,
		appId int,
	) (token string, err error)

	Register(
		context context.Context,
		email string,
		password string,
	) (userId int64, err error)

	IsAdmin(
		context context.Context,
		userId int64,
	) (isAdmin bool, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

const (
	emptyValue = 0
)
