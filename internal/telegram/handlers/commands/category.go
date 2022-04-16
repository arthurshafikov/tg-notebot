package commands

import (
	"errors"
	"fmt"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) HandleAddCategory(message *tgbotapi.Message) error {
	categoryName := message.CommandArguments()

	if err := c.services.Categories.AddCategory(c.ctx, message.Chat.ID, categoryName); err != nil {
		if errors.Is(err, core.ErrCategoryExists) {
			return fmt.Errorf("The category %s already exists!", categoryName)
		}
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Category %s was created successfully!", categoryName))
	_, err := c.bot.Send(msg)

	return err
}

func (c *CommandHandler) HandleRemoveCategory(message *tgbotapi.Message) error {
	return nil
}
