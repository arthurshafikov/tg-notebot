package commands

import (
	"fmt"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) HandleAddNote(message *tgbotapi.Message) error {
	noteContent := message.CommandArguments()

	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Select in which category you want to put this note")

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, category := range categories {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(
			category.Name,
			fmt.Sprintf("%s %s %s", core.AddNoteCommand, category.Name, noteContent),
		)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleRemoveNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Select in which category you want to remove notes")

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, category := range categories {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(
			category.Name,
			fmt.Sprintf("%s %s", core.RemoveNotesChooseCategoryCallbackQuery, category.Name),
		)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleListNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Select in which category you want to list notes")

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, category := range categories {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(
			category.Name,
			fmt.Sprintf("%s %s", core.ListNotesChooseCategoryCallbackQuery, category.Name),
		)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}
