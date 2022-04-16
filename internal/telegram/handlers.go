package telegram

import (
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case core.StartCommand:
		return b.commandHandler.HandleStart(message)
	case core.AddCategoryCommand:
		return b.commandHandler.HandleAddCategory(message)
	case core.RemoveCategoryCommand:
		return b.commandHandler.HandleRemoveCategory(message)
	case core.RenameCategoryCommand:
		return b.commandHandler.HandleRenameCategory(message)
	}

	return nil
}

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Please use one of commands:`) // todo describe commands
	b.bot.Send(msg)

	return nil
}

func (b *TelegramBot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	splittedData := strings.Split(query.Data, " ")
	if len(splittedData) < 2 {
		return core.ErrWrongCallbackQueryData
	}

	switch splittedData[0] {
	case core.RemoveCategoryCommand:
		return b.queryHandler.HandleRemoveCategory(b.ctx, query.Message.Chat.ID, splittedData[1])
	}

	return nil
}
