package services

import (
	"context"
	"errors"
	"strings"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type CategoryService struct {
	repo   repository.Categories
	logger Logger
}

func NewCategoryService(logger Logger, repo repository.Categories) *CategoryService {
	return &CategoryService{
		repo:   repo,
		logger: logger,
	}
}

func (c *CategoryService) AddCategory(ctx context.Context, telegramChatID int64, name string) error {
	if strings.Contains(name, "=") {
		return core.ErrInvalidateCategoryName
	}
	if err := c.repo.AddCategory(ctx, telegramChatID, name); err != nil {
		if !errors.Is(err, core.ErrCategoryExists) {
			c.logger.Error(err)

			return core.ErrServerError
		}

		return core.ErrCategoryExists
	}

	return nil
}

func (c *CategoryService) RemoveCategory(ctx context.Context, telegramChatID int64, name string) error {
	if err := c.repo.RemoveCategory(ctx, telegramChatID, name); err != nil {
		if !errors.Is(err, core.ErrNotFound) {
			c.logger.Error(err)

			return core.ErrServerError
		}

		return core.ErrNotFound
	}

	return nil
}

func (c *CategoryService) RenameCategory(ctx context.Context, telegramChatID int64, name, newName string) error {
	if err := c.repo.RenameCategory(ctx, telegramChatID, name, newName); err != nil {
		if !errors.Is(err, core.ErrNotFound) {
			c.logger.Error(err)

			return core.ErrServerError
		}

		return core.ErrNotFound
	}

	return nil
}

func (c *CategoryService) ListCategories(ctx context.Context, telegramChatID int64) ([]core.Category, error) {
	categories, err := c.repo.ListCategories(ctx, telegramChatID)
	if err != nil {
		if !errors.Is(err, core.ErrNotFound) {
			c.logger.Error(err)

			return categories, core.ErrServerError
		}

		return categories, core.ErrNotFound
	}

	return categories, nil
}
