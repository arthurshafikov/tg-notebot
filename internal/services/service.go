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
	AddNote(ctx context.Context, telegramChatID int64, categoryName, content string) error
	ListNotesFromCategory(ctx context.Context, telegramChatID int64, categoryName string) ([]core.Note, error)
	RemoveNote(ctx context.Context, telegramChatID int64, categoryName, content string) error
}

type Users interface {
	CreateIfNotExists(ctx context.Context, telegramChatID int64) error
	CheckChatIDExists(ctx context.Context, telegramChatID int64) error
}

type Logger interface {
	Error(err error)
}

type Services struct {
	Categories
	Notes
	Users
}

type Deps struct {
	Repository *repository.Repository
	Logger
}

func NewServices(deps Deps) *Services {
	categories := NewCategoryService(deps.Logger, deps.Repository.Categories)
	notes := NewNoteService(deps.Logger, deps.Repository.Notes)
	users := NewUserService(deps.Logger, deps.Repository.Users)

	return &Services{
		Categories: categories,
		Notes:      notes,
		Users:      users,
	}
}
