package queries

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/config"
	"github.com/arthurshafikov/tg-notebot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type QueryHandler struct {
	ctx      context.Context
	bot      *tgbotapi.BotAPI
	services *services.Services
	messages config.Messages
}

func NewQueryHandler(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	services *services.Services,
	messages config.Messages,
) *QueryHandler {
	return &QueryHandler{
		ctx:      ctx,
		bot:      bot,
		services: services,
		messages: messages,
	}
}
