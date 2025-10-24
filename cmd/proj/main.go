package main

import (
"log/slog"
"os"

"github.com/alexshd/projectstarter/internal/cmd"
"github.com/lmittmann/tint"
)

func init() {
	// Initialize structured logging with colored output
	slog.SetDefault(slog.New(
tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelInfo,
			TimeFormat: "15:04:05.0000",
			NoColor:    false,
			AddSource:  false,
		}),
	))
}

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("command failed", "error", err)
		os.Exit(1)
	}
}
