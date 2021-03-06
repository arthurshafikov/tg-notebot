package repository

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
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

type Repository struct {
	Categories
	Notes
	Users
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Categories: mongodb.NewCategory(db),
		Notes:      mongodb.NewNote(db),
		Users:      mongodb.NewUser(db),
	}
}
