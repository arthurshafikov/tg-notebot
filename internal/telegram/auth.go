package telegram

import (
	"github.com/arthurshafikov/tg-notebot/internal/core"
)

func (b *TelegramBot) checkAuthorization(chatID int64) error {
	if err := b.services.Users.CheckChatIDExists(b.ctx, chatID); err != nil {
		return core.ErrNotAuthorized
	}

	return nil
}
