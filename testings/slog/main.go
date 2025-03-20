package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// logger := slog.Default()
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"hello, world",
		slog.String("user", os.Getenv("USER")),
	)
}
