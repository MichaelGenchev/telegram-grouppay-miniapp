package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/MichaelGenchev/telegram-grouppay-miniapp/internal/application"
	"github.com/MichaelGenchev/telegram-grouppay-miniapp/internal/config"
	"github.com/MichaelGenchev/telegram-grouppay-miniapp/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Initialize logger
	log := logger.NewDefault()
	defer log.Sync()

	log.Info("Starting GroupPay Bot", logger.String("version", "1.0.0"))

	cfg := config.Config{
		Logger: logger.Config{
			Level:       getEnvOrDefault("LOG_LEVEL", "info"),
			Environment: getEnvOrDefault("ENVIRONMENT", "development"),
			OutputPath:  getEnvOrDefault("LOG_OUTPUT", "stdout"),
		},
	}

	if v := os.Getenv("TG_BOT_TOKEN"); v != "" {
		cfg.TgBotToken = v
	} else {
		log.Fatal("TG_BOT_TOKEN environment variable is required")
	}

	err := run(ctx, cancel, cfg, log)
	if err != nil {
		log.Error("Application failed", logger.Error(err))
		os.Exit(1)
	}

	log.Info("GroupPay Bot shutdown complete")
}

func run(ctx context.Context, cancel context.CancelFunc, cfg config.Config, log logger.Logger) error {
	app, err := application.New(cfg.TgBotToken, log)
	if err != nil {
		return fmt.Errorf("create application: %w", err)
	}

	log.Info("Application created successfully")
	var wg sync.WaitGroup

	wg.Add(1)
	go app.Run(ctx, &wg)

	<-ctx.Done()
	log.Info("Shutdown signal received, waiting for graceful shutdown...")
	wg.Wait()

	log.Info("Application shutdown complete")
	return nil
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
