package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	startCommand = "/start"
)

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Please use one of commands:`) // todo describe commands
	b.bot.Send(msg)

	return nil
}

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	switch message.Text {
	case startCommand:
		return b.handleStart(message)
	}

	return nil
}

func (b *TelegramBot) handleStart(message *tgbotapi.Message) error {
	if err := b.services.Users.CreateIfNotExists(b.ctx, message.From.UserName, message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, `Successfully authorized!`)
	_, err := b.bot.Send(msg)

	return err
}
