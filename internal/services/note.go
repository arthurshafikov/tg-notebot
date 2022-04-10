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

func (n *NoteService) AddNote(ctx context.Context, userId int64, categoryId int64, note string) error {
	panic("not implemented") // TODO: Implement
}

func (n *NoteService) ListNotes(ctx context.Context, userId int64, categoryId int64) ([]core.Note, error) {
	panic("not implemented") // TODO: Implement
}

func (n *NoteService) RemoveNotes(ctx context.Context, userId []int, notesNumbers []int) error {
	panic("not implemented") // TODO: Implement
}
