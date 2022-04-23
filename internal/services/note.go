package services

import (
	"context"
	"errors"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type NoteService struct {
	repo   repository.Notes
	logger Logger
}

func NewNoteService(logger Logger, repo repository.Notes) *NoteService {
	return &NoteService{
		repo:   repo,
		logger: logger,
	}
}

func (n *NoteService) AddNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	if err := n.repo.AddNote(ctx, telegramChatID, categoryName, content); err != nil {
		n.logger.Error(err)

		return core.ErrServerError
	}

	return nil
}

func (n *NoteService) RemoveNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	if err := n.repo.RemoveNote(ctx, telegramChatID, categoryName, content); err != nil {
		if !errors.Is(err, core.ErrNotFound) {
			n.logger.Error(err)

			return core.ErrServerError
		}

		return core.ErrNotFound
	}

	return nil
}

func (n *NoteService) ListNotesFromCategory(
	ctx context.Context,
	telegramChatID int64,
	categoryName string,
) ([]core.Note, error) {
	notes, err := n.repo.ListNotesFromCategory(ctx, telegramChatID, categoryName)
	if err != nil {
		if !errors.Is(err, core.ErrNotFound) {
			n.logger.Error(err)

			return notes, core.ErrServerError
		}

		return notes, core.ErrNotFound
	}

	return notes, nil
}
