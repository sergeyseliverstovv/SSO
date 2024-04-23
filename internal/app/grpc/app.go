package grpcapp

import (
	//"log"
	"fmt"
	"log/slog"
	"net"

	//"strconv"

	authgrpc "example.com/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic("err")
	}

}

func (a *App) Run() error {
	const op = "grpcappRun"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	log.Info("Starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil

}

// stop stops gRPC server.

func (a *App) Stop() error {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stoping grpc server", slog.Int("port", a.port))

	//log.Info("Stopping gRPC server")

	a.gRPCServer.GracefulStop()

	return nil
}
