package telegram

import (
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case core.StartCommand:
		return b.commandHandler.HandleStart(message)
	// categories
	case core.AddCategoryCommand:
		return b.commandHandler.HandleAddCategory(message)
	case core.RemoveCategoryCommand:
		return b.commandHandler.HandleRemoveCategory(message)
	case core.RenameCategoryCommand:
		return b.commandHandler.HandleRenameCategory(message)
	case core.ListCategoriesCommand:
		return b.commandHandler.HandleListCategories(message)
	// notes
	case core.AddNoteCommand:
		return b.commandHandler.HandleAddNote(message)
	case core.RemoveNotesCommand:
		return b.commandHandler.HandleRemoveNotes(message)
	case core.ListNotes:
		return b.commandHandler.HandleListNotes(message)
	case core.ListAllNotes:
		return b.commandHandler.HandleListAllNotes(message)
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	return b.commandHandler.HandleAddNote(message)
}

func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	splittedData := strings.Split(query.Data, core.SpecialDelimeterInQueryCallback)
	if len(splittedData) < 2 {
		return core.ErrWrongCallbackQueryData
	}

	switch splittedData[0] {
	// categories
	case core.RemoveCategoryCommand:
		return b.queryHandler.HandleRemoveCategory(query.Message.Chat.ID, splittedData[1])
	// notes
	case core.AddNoteCommand:
		return b.queryHandler.HandleAddNote(query.Message.Chat.ID, splittedData[1:])
	case core.RemoveNotesChooseCategoryCallbackQuery:
		return b.queryHandler.HandleListNotesToRemoveInCategory(query.Message.Chat.ID, splittedData[1])
	case core.RemoveNotesCommand:
		return b.queryHandler.HandleRemoveNotes(query.Message.Chat.ID, splittedData[1:])
	case core.ListNotesChooseCategoryCallbackQuery:
		return b.queryHandler.HandleListNotes(query.Message.Chat.ID, splittedData[1])
	}

	return nil
}
