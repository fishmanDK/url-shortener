package main

import (
	"log/slog"
	"net/http"
	"os"
	"test-ozon/config"
	"test-ozon/internal/http-server/handlers"
	"test-ozon/internal/service"
	"test-ozon/internal/storage"

	"github.com/gin-gonic/gin"
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
	service := service.NewService(storage)
	handlers := handlers.NewHandlers(service)
	routs := handlers.InitRouts(logger)

	server := InitServer(cfg, routs)

	server.ListenAndServe()
}

func InitServer(cfg *config.LocalConfig, routs *gin.Engine) *http.Server{
	return &http.Server{
		Addr:         ":8080",
        Handler:      routs,
        ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
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