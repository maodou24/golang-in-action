package _log

import (
	"log/slog"
	"os"
)

func DefaultLogOutput() {
	slog.Info("this is slog", "pid", os.Getpid(), "name", "maodou")
	slog.Debug("debugger", "step", "starting")
	slog.Error("panic", "err", "file not exists")
}

func LogOutput() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(handler).With("version", "1.1.0")

	logger.Info("this is slog", "pid", os.Getpid(), "name", "maodou")
	logger.Debug("debugger", "step", "starting")
	logger.Error("panic", "err", "file not exists")
}
