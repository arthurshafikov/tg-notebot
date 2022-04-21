package commands

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/config"
	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandHandler struct {
	ctx      context.Context
	bot      *tgbotapi.BotAPI
	services *services.Services

	messages config.Messages
}

func NewCommandHandler(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	services *services.Services,
	messages config.Messages,
) *CommandHandler {
	return &CommandHandler{
		ctx:      ctx,
		bot:      bot,
		services: services,
		messages: messages,
	}
}

func (c *CommandHandler) HandleStart(message *tgbotapi.Message) error {
	if err := c.services.Users.CreateIfNotExists(c.ctx, message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.AuthSuccess)

	return c.sendMessage(msg)
}

func (c *CommandHandler) getKeyboardCategories(
	categories []core.Category,
	callable func(category core.Category) (string, string),
) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, category := range categories {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(callable(category))
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	return keyboard
}

func (c *CommandHandler) sendMessage(msg tgbotapi.MessageConfig) error {
	msg.ParseMode = "markdown"
	_, err := c.bot.Send(msg)

	return err
}
