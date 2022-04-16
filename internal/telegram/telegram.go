package telegram

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/services"
	"github.com/arthurshafikov/tg-notebot/internal/telegram/handlers/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	ctx      context.Context
	bot      *tgbotapi.BotAPI
	services *services.Services

	commandHandler *commands.CommandHandler
}

func NewTelegramBot(ctx context.Context, bot *tgbotapi.BotAPI, services *services.Services) *TelegramBot {
	commandHandler := commands.NewCommandHandler(ctx, bot, services)

	return &TelegramBot{
		ctx:      ctx,
		bot:      bot,
		services: services,

		commandHandler: commandHandler,
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
		if err := b.checkAuthorization(update.Message.Chat.ID); err != nil && update.Message.Text != startCommand {
			b.handleError(update.Message.Chat.ID, err)

			continue
		}

		if update.CallbackQuery != nil {
			if err := b.handleCallbackQuery(update.CallbackQuery); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		} else if update.Message == nil {
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
