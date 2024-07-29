package auth

import (
	"context"

	sso "github.com/sergeyseliverstovv/protos/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appId int64) (token string, err error)

	RegisterNewServer(ctx context.Context, email string, password string) (userId int64, err error)

	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type serverApi struct {
	sso.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	sso.RegisterAuthServer(gRPC, &serverApi{auth: auth})
}

func (s *serverApi) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int64(req.GetAppId()))

	if err != nil {
		// TODO ...
		return nil, status.Error(codes.Internal, "internal error")
	}
	// TODO: Implement actual login logic using provided email, password, and app_id.

	return &sso.LoginResponse{Token: token}, nil
}

func (s *serverApi) Register(ctx context.Context, req *sso.RigisterRequest) (*sso.RigisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
	userId, err := s.auth.RegisterNewServer(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &sso.RigisterResponse{UserId: userId}, nil
}

func (s *serverApi) IsAdmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	if err := validaeIsAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, int64(req.GetUserId()))

	if err != nil {
		// TODO:...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &sso.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func validateRegister(reg *sso.RigisterRequest) error {
	if reg.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if reg.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func validaeIsAdmin(reg *sso.IsAdminRequest) error {
	if reg.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	return nil
}
