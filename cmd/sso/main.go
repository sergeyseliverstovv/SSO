package main

import (
	"log/slog"
	"os"

	"example.com/sso/internal/config"
	slogpretty "example.com/sso/internal/lib/logger/handlers/slogprrety"
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

	// TODO: инициализировать логгер

	// TODO: инициализация приложения(app)

	// TODO: запустить gRPC сервер приложения

}

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
