package services

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/core"
)

type Categories interface {
	AddCategory(ctx context.Context, userId int64, name string) error
	RemoveCategory(ctx context.Context, userId, id int64) error
	RenameCategory(ctx context.Context, userId, id int64, newName string) error
	ListCategories(ctx context.Context, userId int64) ([]core.Category, error)
}

type Notes interface {
	AddNote(ctx context.Context, userId, categoryId int64, note string) error
	ListNotes(ctx context.Context, userId, categoryId int64) ([]core.Note, error)
	RemoveNotes(ctx context.Context, userId, notesNumbers []int) error
}

type Services struct {
	Categories
	Notes
}

type Deps struct{}

func NewServices(deps Deps) *Services {
	return &Services{
		Categories: NewCategoryService(),
		Notes:      NewNoteService(),
	}
}
