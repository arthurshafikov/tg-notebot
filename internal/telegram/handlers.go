package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	startCommand       = "start"
	addCategoryCommand = "addcategory"
)

func (b *TelegramBot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Please use one of commands:`) // todo describe commands
	b.bot.Send(msg)

	return nil
}

func (b *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStart(message)
	case addCategoryCommand:
		return b.handleAddCategory(message)
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

func (b *TelegramBot) handleAddCategory(message *tgbotapi.Message) error {
	categoryName := message.CommandArguments()

	if err := b.services.Categories.AddCategory(b.ctx, message.Chat.ID, categoryName); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Category %s was created successfully!", categoryName))
	_, err := b.bot.Send(msg)

	return err
}
