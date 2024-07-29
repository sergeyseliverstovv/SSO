package app

import (
	"log/slog"
	"time"

	grpcapp "sso/internal/app/grpc"
	//"example.com/sso/internal/grpc/auth"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {

	//TODO: инициализировать хранилище storage

	//TODO: инициализировать auth service

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
