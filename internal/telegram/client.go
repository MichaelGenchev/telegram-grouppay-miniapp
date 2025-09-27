package telegram

import (
	"context"
	"fmt"

	"github.com/MichaelGenchev/telegram-grouppay-miniapp/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Client wraps the telegram bot client with our configuration
type Client struct {
	bot    *bot.Bot
	token  string
	logger logger.Logger
}

// New creates a new telegram client
func New(token string, log logger.Logger) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("bot token is required")
	}

	client := &Client{
		token:  token,
		logger: log.With(logger.String("component", "telegram")),
	}

	b, err := bot.New(token, bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Default handler for unhandled updates
		client.logger.Debug("Unhandled update received",
			logger.Any("update", update),
			logger.Int64("chat_id", update.Message.Chat.ID),
		)
	}))
	if err != nil {
		return nil, fmt.Errorf("create bot client: %w", err)
	}

	client.bot = b
	client.logger.Info("Telegram client initialized successfully")

	return client, nil
}

// Start starts the bot with the given context
func (c *Client) Start(ctx context.Context) {
	c.logger.Info("Starting Telegram bot")
	c.bot.Start(ctx)
	c.logger.Info("Telegram bot stopped")
}

// RegisterHandler registers a command handler with the bot
func (c *Client) RegisterHandler(handlerType bot.HandlerType, pattern string, matchType bot.MatchType, handler bot.HandlerFunc) {
	c.bot.RegisterHandler(handlerType, pattern, matchType, handler)
}

// Bot returns the underlying bot instance for direct access when needed
func (c *Client) Bot() *bot.Bot {
	return c.bot
}
