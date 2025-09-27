package handlers

import (
	"context"

	"github.com/MichaelGenchev/telegram-grouppay-miniapp/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// CommandHandler handles telegram bot commands
type CommandHandler struct {
	logger logger.Logger
	// Add dependencies here (database, services, etc.)
}

// New creates a new command handler
func New(log logger.Logger) *CommandHandler {
	return &CommandHandler{
		logger: log.With(logger.String("component", "handlers")),
	}
}

// RegisterHandlers registers all command handlers with the telegram client
func (h *CommandHandler) RegisterHandlers(registerFunc func(handlerType bot.HandlerType, pattern string, matchType bot.MatchType, handler bot.HandlerFunc)) {
	h.logger.Info("Registering command handlers")

	// Register all command handlers
	registerFunc(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.HandleStart)
	registerFunc(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, h.HandleHelp)
	registerFunc(bot.HandlerTypeMessageText, "/create_group", bot.MatchTypeExact, h.HandleCreateGroup)
	registerFunc(bot.HandlerTypeMessageText, "/add_expense", bot.MatchTypeExact, h.HandleAddExpense)
	registerFunc(bot.HandlerTypeMessageText, "/balance", bot.MatchTypeExact, h.HandleBalance)
	registerFunc(bot.HandlerTypeMessageText, "/settle", bot.MatchTypeExact, h.HandleSettle)

	h.logger.Info("Command handlers registered successfully")
}

// HandleStart handles the /start command
func (h *CommandHandler) HandleStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /start command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.String("username", update.Message.From.Username),
		logger.Int64("chat_id", update.Message.Chat.ID),
		logger.String("command", "/start"),
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to GroupPay! 🎉\n\nI'll help you manage shared expenses with your friends.\n\nUse /help to see available commands.",
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send start message",
			logger.Error(err),
			logger.Int64("chat_id", update.Message.Chat.ID),
		)
	} else {
		h.logger.DebugContext(ctx, "Start message sent successfully",
			logger.Int64("chat_id", update.Message.Chat.ID),
		)
	}
}

// HandleHelp handles the /help command
func (h *CommandHandler) HandleHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /help command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.Int64("chat_id", update.Message.Chat.ID),
	)

	helpText := `Available commands:

/start - Welcome message
/help - Show this help
/create_group - Create a new expense group
/add_expense - Add an expense
/balance - Show current balances
/settle - Settle up expenses`

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   helpText,
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send help message", logger.Error(err))
	}
}

// HandleCreateGroup handles the /create_group command
func (h *CommandHandler) HandleCreateGroup(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /create_group command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.Int64("chat_id", update.Message.Chat.ID),
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Creating expense groups is coming soon! 🚧",
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send create group message", logger.Error(err))
	}
}

// HandleAddExpense handles the /add_expense command
func (h *CommandHandler) HandleAddExpense(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /add_expense command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.Int64("chat_id", update.Message.Chat.ID),
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Adding expenses is coming soon! 💰",
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send add expense message", logger.Error(err))
	}
}

// HandleBalance handles the /balance command
func (h *CommandHandler) HandleBalance(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /balance command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.Int64("chat_id", update.Message.Chat.ID),
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Balance checking is coming soon! 📊",
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send balance message", logger.Error(err))
	}
}

// HandleSettle handles the /settle command
func (h *CommandHandler) HandleSettle(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.logger.InfoContext(ctx, "Received /settle command",
		logger.Int64("user_id", update.Message.From.ID),
		logger.Int64("chat_id", update.Message.Chat.ID),
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Settlement is coming soon! 🤝",
	})
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to send settle message", logger.Error(err))
	}
}
