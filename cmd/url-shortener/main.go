package main

import (
	"log/slog"
	"os"
	"test-ozon/config"
	"test-ozon/internal/http-server/handlers"
	"test-ozon/internal/service/api"
	"test-ozon/internal/storage"
)

const (
	envLocal = "local"
	envDev = "dev"

	dbMemory = "memory"
	dbPostgres = "postgres"
)

func main() {

	cfg := config.NewLocalConfig()

	logger := setupLogger(cfg.Env)

	storage, err := storage.NewDB(dbPostgres)
	if err != nil{
		logger.Debug(err.Error())
		return
	}
	service := api.NewServiceApi(storage)
	handlers := handlers.NewHandlers(service)

	server := handlers.InitRouts(logger)

	// server.ListenAndServe()
	server.Run(cfg.Server.Address)
}

func setupLogger(env string) *slog.Logger{
	var logger *slog.Logger

	switch env{
	case envLocal:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slogHandler := slog.NewTextHandler(os.Stdout, opts)
		logger = slog.New(slogHandler)
	case envDev:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slogHandler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(slogHandler)
	}

	return logger
}