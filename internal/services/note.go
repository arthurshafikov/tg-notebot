package services

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type NoteService struct {
	repo repository.Notes
}

func NewNoteService(repo repository.Notes) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (n *NoteService) AddNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	return n.repo.AddNote(ctx, telegramChatID, categoryName, content)
}

func (n *NoteService) ListNotesFromCategory(ctx context.Context, telegramChatID int64, categoryName string) ([]core.Note, error) {
	return n.repo.ListNotesFromCategory(ctx, telegramChatID, categoryName)
}

func (n *NoteService) RemoveNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	return n.repo.RemoveNote(ctx, telegramChatID, categoryName, content)
}
