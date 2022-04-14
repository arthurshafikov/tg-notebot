package services

import (
	"context"

	"github.com/arthurshafikov/tg-notebot/internal/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateIfNotExists(ctx context.Context, userName string, telegramChatID int64) error {
	return u.repo.CreateIfNotExists(ctx, userName, telegramChatID)
}
