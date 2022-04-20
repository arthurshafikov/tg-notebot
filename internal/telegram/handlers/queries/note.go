package queries

import (
	"fmt"
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (q *QueryHandler) HandleAddNote(telegramChatID int64, args []string) error {
	if len(args) < 2 {
		return core.ErrServerError
	}

	categoryName := args[0]
	noteContent := strings.Join(args[1:], " ")

	if err := q.services.Notes.AddNote(q.ctx, telegramChatID, categoryName, noteContent); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, q.messages.NoteCreated)

	return q.sendMessage(msg)
}

func (q *QueryHandler) HandleListNotesToRemoveInCategory(telegramChatID int64, categoryName string) error {
	notes, err := q.services.Notes.ListNotesFromCategory(q.ctx, telegramChatID, categoryName)
	if err != nil {
		return err
	}

	var msg tgbotapi.MessageConfig
	if len(notes) > 0 {
		msg = tgbotapi.NewMessage(telegramChatID, q.messages.SelectNotes)

		keyboard := tgbotapi.InlineKeyboardMarkup{}
		for _, note := range notes {
			var row []tgbotapi.InlineKeyboardButton
			btn := tgbotapi.NewInlineKeyboardButtonData(
				note.Content,
				fmt.Sprintf("%s %s %s", core.RemoveNotesCommand, categoryName, note.Content),
			)
			row = append(row, btn)
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
		}

		msg.ReplyMarkup = keyboard
	} else {
		msg = tgbotapi.NewMessage(telegramChatID, fmt.Sprintf(q.messages.NoNotesInCategory, categoryName))
	}

	return q.sendMessage(msg)
}

func (q *QueryHandler) HandleRemoveNotes(telegramChatID int64, args []string) error {
	if len(args) < 2 {
		return core.ErrServerError
	}

	categoryName := args[0]
	noteContent := strings.Join(args[1:], " ")

	if err := q.services.Notes.RemoveNote(q.ctx, telegramChatID, categoryName, noteContent); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, q.messages.NoteRemoved)

	return q.sendMessage(msg)
}

func (q *QueryHandler) HandleListNotes(telegramChatID int64, categoryName string) error {
	notes, err := q.services.Notes.ListNotesFromCategory(q.ctx, telegramChatID, categoryName)
	if err != nil {
		return err
	}

	var msgText string
	if len(notes) > 0 {
		msgText = fmt.Sprintf(q.messages.ListNotesInCategory, categoryName)
		for _, note := range notes {
			msgText += fmt.Sprintf("\n - %s", note.Content)
		}
	} else {
		msgText = fmt.Sprintf(q.messages.NoNotesInCategory, categoryName)
	}

	msg := tgbotapi.NewMessage(telegramChatID, msgText)

	return q.sendMessage(msg)
}
