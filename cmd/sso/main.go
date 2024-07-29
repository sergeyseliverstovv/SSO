package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"sso/internal/app"
	"sso/internal/config"
	slogpretty "sso/internal/lib/logger/handlers/slogprrety"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port))

	log.Debug("debug messange")

	log.Error("error messange")

	log.Warn("warn messange")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTl)

	go application.GRPCSrv.MustRun() // Запускаем gRPC сервер в отдельной go - рутине

	// TODO: запустить gRPC сервер приложения

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) // Слушаем сигналы из os SIGTERM и SIGINT

	sign := <-stop // Ожидаем когда что то придет в канал

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop() // Останавливаем приложение

	log.Info("application stopped")

}

// Определяет в каком типе будут выводится логи в зависимости от того в каком окружение запущен сервис
func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log

}

func setupPrettySlog() *slog.Logger {

	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
