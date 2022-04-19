package telegram

import "fmt"

func (b *TelegramBot) checkAuthorization(chatID int64) error {
	if err := b.services.Users.CheckChatIDExists(b.ctx, chatID); err != nil {
		return fmt.Errorf(b.messages.NotAuthorized)
	}

	return nil
}
