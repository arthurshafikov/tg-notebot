package services

import (
	"context"
	"errors"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type UserService struct {
	repo   repository.Users
	logger Logger
}

func NewUserService(logger Logger, repo repository.Users) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) CreateIfNotExists(ctx context.Context, telegramChatID int64) error {
	if err := u.repo.CreateIfNotExists(ctx, telegramChatID); err != nil {
		u.logger.Error(err)

		return core.ErrServerError
	}

	return nil
}

func (u *UserService) CheckChatIDExists(ctx context.Context, telegramChatID int64) error {
	if err := u.repo.CheckChatIDExists(ctx, telegramChatID); err != nil {
		if !errors.Is(core.ErrNotFound, err) {
			u.logger.Error(err)

			return core.ErrServerError
		}

		return core.ErrNotFound
	}

	return nil
}
