package services

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/core"
	"github.com/thewolf27/wolf-notebot/internal/repository"
)

type CategoryService struct {
	repo repository.Categories
}

func NewCategoryService(repo repository.Categories) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) AddCategory(ctx context.Context, userId int64, name string) error {
	panic("not implemented") // TODO: Implement
}

func (c *CategoryService) RemoveCategory(ctx context.Context, userId int64, id int64) error {
	panic("not implemented") // TODO: Implement
}

func (c *CategoryService) RenameCategory(ctx context.Context, userId int64, id int64, newName string) error {
	panic("not implemented") // TODO: Implement
}

func (c *CategoryService) ListCategories(ctx context.Context, userId int64) ([]core.Category, error) {
	panic("not implemented") // TODO: Implement
}
