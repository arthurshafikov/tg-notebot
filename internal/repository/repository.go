package repository

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/core"
	"github.com/thewolf27/wolf-notebot/internal/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Categories interface {
	AddCategory(ctx context.Context, userId int, name string) error
	RemoveCategory(ctx context.Context, userId int, name string) error
	RenameCategory(ctx context.Context, userId int, name, newName string) error
	ListCategories(ctx context.Context, userId int) ([]core.Category, error)
}

type Notes interface {
	AddNote(ctx context.Context, userId, categoryId int64, note string) error
	ListNotes(ctx context.Context, userId, categoryId int64) ([]core.Note, error)
	RemoveNotes(ctx context.Context, userId, notesNumbers []int) error
}

type Repository struct {
	Categories
	Notes
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Categories: mongodb.NewCategory(db),
	}
}
