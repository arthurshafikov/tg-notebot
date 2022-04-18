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

	msg := tgbotapi.NewMessage(telegramChatID, "Note was added successfully!")
	_, err := q.bot.Send(msg)

	return err
}

func (q *QueryHandler) HandleListNotesToRemoveInCategory(telegramChatID int64, categoryName string) error {
	notes, err := q.services.Notes.ListNotesFromCategory(q.ctx, telegramChatID, categoryName)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, "Select notes to remove")

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
	_, err = q.bot.Send(msg)

	return err
}

func (q *QueryHandler) HandleRemoveNotes(telegramChatID int64, args []string) error {
	if len(args) < 2 {
		return core.ErrServerError
	}

	categoryName := args[0]
	noteContent := strings.Join(args[1:], " ")

	// remove unessesary context?
	if err := q.services.Notes.RemoveNote(q.ctx, telegramChatID, categoryName, noteContent); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(telegramChatID, "Note was removed successfully!")
	_, err := q.bot.Send(msg)

	return err
}

func (q *QueryHandler) HandleListNotes(telegramChatID int64, categoryName string) error {
	notes, err := q.services.Notes.ListNotesFromCategory(q.ctx, telegramChatID, categoryName)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf("Here is your notes in category %s:", categoryName)
	for _, note := range notes {
		msgText += fmt.Sprintf("\n - %s", note.Content)
	}

	msg := tgbotapi.NewMessage(telegramChatID, msgText)
	_, err = q.bot.Send(msg)

	return err
}
