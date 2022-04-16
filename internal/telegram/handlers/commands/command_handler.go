package commands

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandHandler struct {
	ctx      context.Context
	bot      *tgbotapi.BotAPI
	services *services.Services
}

func NewCommandHandler(ctx context.Context, bot *tgbotapi.BotAPI, services *services.Services) *CommandHandler {
	return &CommandHandler{
		ctx:      ctx,
		bot:      bot,
		services: services,
	}
}

func (c *CommandHandler) HandleStart(message *tgbotapi.Message) error {
	if err := c.services.Users.CreateIfNotExists(c.ctx, message.From.UserName, message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, `Successfully authorized!`)
	_, err := c.bot.Send(msg)

	return err
}
