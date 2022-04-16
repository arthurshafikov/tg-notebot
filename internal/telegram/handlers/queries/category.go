package queries

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (q *QueryHandler) HandleRemoveCategory(ctx context.Context, telegramChatID int64, name string) error {
	if err := q.services.Categories.RemoveCategory(q.ctx, telegramChatID, name); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, fmt.Sprintf("Category %s was successfully removed!", name))
	_, err := q.bot.Send(msg)

	return err
}
