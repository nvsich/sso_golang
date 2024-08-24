package auth

import (
	"context"
	ssov1 "github.com/nvsich/sso_protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := validateLoginRequest(in); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func validateLoginRequest(in *ssov1.LoginRequest) error {
	if in.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}

	if in.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}

	if in.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "missing app_id")
	}

	return nil
}
