package queries

import (
	"context"
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (q *QueryHandler) HandleAddNote(ctx context.Context, telegramChatID int64, args []string) error {
	if len(args) < 2 {
		return core.ErrServerError
	}

	categoryName := args[0]
	noteContent := strings.Join(args[1:], " ")

	if err := q.services.Notes.AddNote(q.ctx, telegramChatID, categoryName, noteContent); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, "Note was added successfully!")
	_, err := q.bot.Send(msg)

	return err
}
