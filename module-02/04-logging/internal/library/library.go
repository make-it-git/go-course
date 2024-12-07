package library

import "log/slog"

func Call(logger *slog.Logger) {
	logger.Info("Hi, I'm library")
}
