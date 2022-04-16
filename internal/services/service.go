package services

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type Categories interface {
	AddCategory(ctx context.Context, telegramChatID int64, name string) error
	RemoveCategory(ctx context.Context, telegramChatID int64, name string) error
	RenameCategory(ctx context.Context, telegramChatID int64, name, newName string) error
	ListCategories(ctx context.Context, telegramChatID int64) ([]core.Category, error)
}

type Notes interface {
	AddNote(ctx context.Context, userId, categoryId int64, note string) error
	ListNotes(ctx context.Context, userId, categoryId int64) ([]core.Note, error)
	RemoveNotes(ctx context.Context, userId, notesNumbers []int) error
}

type Users interface {
	CreateIfNotExists(ctx context.Context, userName string, telegramChatID int64) error
	CheckChatIDExists(ctx context.Context, telegramChatID int64) error
}

type Services struct {
	Categories
	Notes
	Users
}

type Deps struct {
	Repository *repository.Repository
}

func NewServices(deps Deps) *Services {
	return &Services{
		Categories: NewCategoryService(deps.Repository.Categories),
		Notes:      NewNoteService(deps.Repository.Notes),
		Users:      NewUserService(deps.Repository.Users),
	}
}
