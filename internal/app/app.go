package app

import (
	"log/slog"
	"time"

	grpcapp "example.com/sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePatn string, tokenTTL time.Duration) *App {

	//TODO: инициализировать хранилище storage

	//TODO: инициализировать auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
