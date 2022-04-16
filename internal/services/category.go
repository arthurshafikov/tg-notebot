package services

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type CategoryService struct {
	repo repository.Categories
}

func NewCategoryService(repo repository.Categories) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) AddCategory(ctx context.Context, telegramChatID int64, name string) error {
	return c.repo.AddCategory(ctx, telegramChatID, name)
}

func (c *CategoryService) RemoveCategory(ctx context.Context, telegramChatID int64, name string) error {
	return c.repo.RemoveCategory(ctx, telegramChatID, name)
}

func (c *CategoryService) RenameCategory(ctx context.Context, telegramChatID int64, name, newName string) error {
	return c.repo.RenameCategory(ctx, telegramChatID, name, newName)
}

func (c *CategoryService) ListCategories(ctx context.Context, telegramChatID int64) ([]core.Category, error) {
	return c.repo.ListCategories(ctx, telegramChatID)
}
