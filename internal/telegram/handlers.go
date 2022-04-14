package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Please use one of commands:`)
	b.bot.Send(msg)

	return nil
}

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	return nil
}
