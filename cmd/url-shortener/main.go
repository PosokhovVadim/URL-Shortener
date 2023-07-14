package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//DONE: Implement config: cleanenv

	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)
	//TODO: Implement logger: slog

	slog.SetDefault(SetupLogger(cfg.Env).With("env", cfg.Env))
	slog.Info("Starting url-shortener")
	slog.Debug("Starting debug logger")

	//TODO: Implement storage: mongodb 

	//TODO: Implement router: chi, "chi render"

	//TODO: To run server:

}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
