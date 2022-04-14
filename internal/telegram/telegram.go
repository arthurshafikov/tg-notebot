package telegram

import (
	"github.com/arthurshafikov/tg-notebot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	bot      *tgbotapi.BotAPI
	services *services.Services
}

func NewTelegramBot(bot *tgbotapi.BotAPI, services *services.Services) *TelegramBot {
	return &TelegramBot{
		bot:      bot,
		services: services,
	}
}

func (b *TelegramBot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
