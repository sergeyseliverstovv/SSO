package auth

import (
	"context"

	sso "github.com/sergeyseliverstovv/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverApi struct {
	sso.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sso.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	panic("implement me")
}

func (s *serverApi) Register(ctx context.Context, req *sso.RigisterRequest) (*sso.RigisterResponse, error) {
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	panic("implement me")
}
