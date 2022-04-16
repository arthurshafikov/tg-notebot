package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	startCommand          = "start"
	addCategoryCommand    = "addcategory"
	removeCategoryCommand = "removecategory"
)

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.commandHandler.HandleStart(message)
	case addCategoryCommand:
		return b.commandHandler.HandleAddCategory(message)
	case removeCategoryCommand:
		return b.commandHandler.HandleRemoveCategory(message)
	}

	return nil
}

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Please use one of commands:`) // todo describe commands
	b.bot.Send(msg)

	return nil
}

func (b *TelegramBot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Ok, I remember")
	b.bot.Send(msg)

	return nil
}
