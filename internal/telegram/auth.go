package telegram

import (
	"errors"
	"fmt"

	"github.com/arthurshafikov/tg-notebot/internal/core"
)

func (b *TelegramBot) checkAuthorization(chatID int64) error {
	if err := b.services.Users.CheckChatIDExists(b.ctx, chatID); err != nil {
		if errors.Is(core.ErrNotFound, err) {
			return fmt.Errorf(b.messages.NotAuthorized)
		}

		return err
	}

	return nil
}
