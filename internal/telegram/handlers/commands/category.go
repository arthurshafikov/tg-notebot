package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) HandleAddCategory(message *tgbotapi.Message) error {
	categoryName := message.CommandArguments()

	if err := c.services.Categories.AddCategory(c.ctx, message.Chat.ID, categoryName); err != nil {
		if errors.Is(err, core.ErrCategoryExists) {
			return fmt.Errorf(c.messages.CategoryExists, categoryName)
		}
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(c.messages.CategoryCreated, categoryName))
	_, err := c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleRemoveCategory(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, c.messages.SelectCategories)

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, category := range categories {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(
			category.Name,
			fmt.Sprintf("%s %s", core.RemoveCategoryCommand, category.Name),
		)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	_, err = c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleRenameCategory(message *tgbotapi.Message) error {
	args := strings.Split(message.CommandArguments(), " ")
	if len(args) != 2 {
		return fmt.Errorf(c.messages.RenameCategoryWrongSyntax)
	}

	if err := c.services.Categories.RenameCategory(c.ctx, message.Chat.ID, args[0], args[1]); err != nil {
		if errors.Is(err, core.ErrCategoryExists) {
			return fmt.Errorf(c.messages.CategoryExists, args[1])
		} else if errors.Is(err, core.ErrNotFound) {
			return fmt.Errorf(c.messages.CategoryNotFound, args[0])
		}

		return err
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf(c.messages.CategoryRenamed, args[0], args[1]),
	)
	_, err := c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleListCategories(message *tgbotapi.Message) error {
	categories, err := c.services.Categories.ListCategories(c.ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msgText := c.messages.ListCategories
	for _, category := range categories {
		msgText += fmt.Sprintf("\n - %s [%v]", category.Name, len(category.Notes))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err = c.bot.Send(msg)

	return err
}
