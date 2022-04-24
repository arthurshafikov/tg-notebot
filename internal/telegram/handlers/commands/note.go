package commands

import (
	"errors"
	"fmt"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) HandleAddNote(message *tgbotapi.Message) error {
	noteContent := message.CommandArguments()
	if noteContent == "" {
		noteContent = message.Text
	}

	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		return errors.New(c.messages.NoCategories)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryForNote)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) []string {
		return []string{core.AddNoteCommand, category.Name, noteContent}
	})

	msg.ReplyMarkup = keyboard

	return c.sendMessage(msg)
}

func (c *CommandHandler) HandleRemoveNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		return fmt.Errorf(c.messages.NoCategories)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryToRemoveNotes)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) []string {
		return []string{core.RemoveNotesChooseCategoryCallbackQuery, category.Name}
	})

	msg.ReplyMarkup = keyboard

	return c.sendMessage(msg)
}

func (c *CommandHandler) HandleListNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		return fmt.Errorf(c.messages.NoCategories)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategoryToListNotes)

	keyboard := c.getKeyboardCategories(categories, func(category core.Category) []string {
		return []string{core.ListNotesChooseCategoryCallbackQuery, category.Name}
	})

	msg.ReplyMarkup = keyboard

	return c.sendMessage(msg)
}

func (c *CommandHandler) HandleListAllNotes(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		return fmt.Errorf(c.messages.NoCategories)
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

	return c.sendMessage(msg)
}
