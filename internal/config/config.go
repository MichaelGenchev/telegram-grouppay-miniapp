package config

import "github.com/MichaelGenchev/telegram-grouppay-miniapp/pkg/logger"

type Config struct {
	TgBotToken string
	Logger     logger.Config
}
