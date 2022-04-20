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

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryForNote)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) (string, string) {
		return category.Name, fmt.Sprintf("%s %s %s", core.AddNoteCommand, category.Name, noteContent)
	})

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleRemoveNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryToRemoveNotes)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) (string, string) {
		return category.Name, fmt.Sprintf("%s %s", core.RemoveNotesChooseCategoryCallbackQuery, category.Name)
	})

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleListNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryToListNotes)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) (string, string) {
		return category.Name, fmt.Sprintf("%s %s", core.ListNotesChooseCategoryCallbackQuery, category.Name)
	})

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleListAllNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msgText := c.messages.ListNotes

	for _, category := range categories {
		msgText += fmt.Sprintf("\n %s:", category.Name)
		for _, note := range category.Notes {
			msgText += fmt.Sprintf("\n - %s", note.Content)
		}
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err = c.bot.Send(msg)

	return err
}
