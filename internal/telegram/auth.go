package telegram

import (
	"errors"
	"fmt"

	"github.com/arthurshafikov/tg-notebot/internal/core"
)

func (b *Bot) checkAuthorization(chatID int64) error {
	if err := b.services.Users.CheckChatIDExists(b.ctx, chatID); err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return fmt.Errorf(b.messages.NotAuthorized)
		}

		return err
	}

	return nil
}
