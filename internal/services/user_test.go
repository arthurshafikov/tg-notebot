package services

import (
	"context"
	"testing"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	mock_repository "github.com/arthurshafikov/tg-notebot/internal/repository/mocks"
	mock_services "github.com/arthurshafikov/tg-notebot/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func getUserRepositoryAndLogger(
	t *testing.T,
) (context.Context, *mock_repository.MockUsers, *mock_services.MockLogger) {
	t.Helper()
	ctrl := gomock.NewController(t)

	return context.Background(), mock_repository.NewMockUsers(ctrl), mock_services.NewMockLogger(ctrl)
}

func TestCreateIfNotExists(t *testing.T) {
	ctx, repo, logger := getUserRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().CreateIfNotExists(ctx, telegramChatID).Times(1).Return(nil),
	)
	service := NewUserService(logger, repo)

	err := service.CreateIfNotExists(ctx, telegramChatID)
	require.NoError(t, err)
}

func TestCreateIfNotExistsReturnServerError(t *testing.T) {
	ctx, repo, logger := getUserRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().CreateIfNotExists(ctx, telegramChatID).Times(1).Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewUserService(logger, repo)

	err := service.CreateIfNotExists(ctx, telegramChatID)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestCheckChatIDExists(t *testing.T) {
	ctx, repo, logger := getUserRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().CheckChatIDExists(ctx, telegramChatID).Times(1).Return(nil),
	)
	service := NewUserService(logger, repo)

	err := service.CheckChatIDExists(ctx, telegramChatID)
	require.NoError(t, err)
}

func TestCheckChatIDExistsReturnServerError(t *testing.T) {
	ctx, repo, logger := getUserRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().CheckChatIDExists(ctx, telegramChatID).Times(1).Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewUserService(logger, repo)

	err := service.CheckChatIDExists(ctx, telegramChatID)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestCheckChatIDExistsUserNotFound(t *testing.T) {
	ctx, repo, logger := getUserRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().CheckChatIDExists(ctx, telegramChatID).Times(1).Return(core.ErrNotFound),
	)
	service := NewUserService(logger, repo)

	err := service.CheckChatIDExists(ctx, telegramChatID)
	require.ErrorIs(t, err, core.ErrNotFound)
}
