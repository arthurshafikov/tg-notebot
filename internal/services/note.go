package services

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/core"
)

type NoteService struct {
}

func NewNoteService() *NoteService {
	return &NoteService{}
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
