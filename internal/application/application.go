package application

import (
	"context"
	"fmt"
	"sync"

	"github.com/MichaelGenchev/telegram-grouppay-miniapp/internal/handlers"
	"github.com/MichaelGenchev/telegram-grouppay-miniapp/internal/telegram"
	"github.com/MichaelGenchev/telegram-grouppay-miniapp/pkg/logger"
)

// Application represents the main application
type Application struct {
	telegramClient *telegram.Client
	commandHandler *handlers.CommandHandler
	logger         logger.Logger
}

// New creates a new application instance
func New(botToken string, log logger.Logger) (*Application, error) {
	log.Info("Initializing application", logger.String("component", "application"))

	// Create telegram client
	telegramClient, err := telegram.New(botToken, log)
	if err != nil {
		return nil, fmt.Errorf("create telegram client: %w", err)
	}

	// Create command handler
	commandHandler := handlers.New(log)

	// Register handlers with telegram client
	commandHandler.RegisterHandlers(telegramClient.RegisterHandler)

	log.Info("Application initialized successfully")

	return &Application{
		telegramClient: telegramClient,
		commandHandler: commandHandler,
		logger:         log,
	}, nil
}

// Run starts the application
func (app *Application) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer app.logger.Sync() // Ensure logs are flushed

	app.logger.Info("Starting application")

	// Start the telegram bot
	app.telegramClient.Start(ctx)

	app.logger.Info("Application stopped")
}
