package grpcapp

import (
	//"log"
	"fmt"
	"log/slog"
	"net"

	//"strconv"

	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, authService authgrpc.Auth, port int) *App {
	gRPCServer := grpc.NewServer()

	//auth := authgrpc.App

	authgrpc.Register(gRPCServer, authService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

// Функция MustRun является оберткой для функции Run. Если не запустился сервис т.е Run вернула ошибку мы бросаем панику.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic("err")
	}

}

func (a *App) Run() error {
	const op = "grpcapp.Run"

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
