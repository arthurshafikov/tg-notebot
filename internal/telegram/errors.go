package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *TelegramBot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, err.Error())
	b.bot.Send(msg)
}
