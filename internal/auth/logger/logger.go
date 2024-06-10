package logger

import (
	"fmt"
	"io"
	"log/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func New(env string, logStorage io.Writer) *slog.Logger {
	return initializeLogger(env, logStorage)
}

func initializeLogger(env string, storage io.Writer) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(storage, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewTextHandler(storage, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(storage, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic(fmt.Sprintf("Unkown environment %s. Set env in config equal to some `local, dev, prod`", env))
	}
	return logger
}
