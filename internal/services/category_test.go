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

var (
	telegramChatID  int64  = int64(20)
	categoryName    string = "newCategoryName"
	categoryNewName string = "categoryNewName"
)

func getCategoryRepositoryAndLogger(
	t *testing.T,
) (context.Context, *mock_repository.MockCategories, *mock_services.MockLogger) {
	t.Helper()
	ctrl := gomock.NewController(t)

	return context.Background(), mock_repository.NewMockCategories(ctrl), mock_services.NewMockLogger(ctrl)
}

func TestAddCategory(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().AddCategory(ctx, telegramChatID, categoryName).Times(1).Return(nil),
	)
	service := NewCategoryService(logger, repo)

	err := service.AddCategory(ctx, telegramChatID, categoryName)
	require.NoError(t, err)
}

func TestAddCategoryCausedServerError(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().AddCategory(ctx, telegramChatID, categoryName).Times(1).Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewCategoryService(logger, repo)

	err := service.AddCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestAddCategoryCategoryExists(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().AddCategory(ctx, telegramChatID, categoryName).Times(1).Return(core.ErrCategoryExists),
		logger.EXPECT().Error(core.ErrCategoryExists).Times(0),
	)
	service := NewCategoryService(logger, repo)

	err := service.AddCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrCategoryExists)
}

func TestRemoveCategory(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().RemoveCategory(ctx, telegramChatID, categoryName).Times(1).Return(nil),
	)
	service := NewCategoryService(logger, repo)

	err := service.RemoveCategory(ctx, telegramChatID, categoryName)
	require.NoError(t, err)
}

func TestRemoveCategoryCausedServerError(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().RemoveCategory(ctx, telegramChatID, categoryName).Times(1).Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewCategoryService(logger, repo)

	err := service.RemoveCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestRemoveCategoryCategoryDoesNotExists(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().RemoveCategory(ctx, telegramChatID, categoryName).Times(1).Return(core.ErrNotFound),
		logger.EXPECT().Error(core.ErrNotFound).Times(0),
	)
	service := NewCategoryService(logger, repo)

	err := service.RemoveCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrNotFound)
}

func TestRenameCategory(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().RenameCategory(ctx, telegramChatID, categoryName, categoryNewName).Times(1).Return(nil),
	)
	service := NewCategoryService(logger, repo)

	err := service.RenameCategory(ctx, telegramChatID, categoryName, categoryNewName)
	require.NoError(t, err)
}

func TestRenameCategoryCausedServerError(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().
			RenameCategory(ctx, telegramChatID, categoryName, categoryNewName).
			Times(1).
			Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewCategoryService(logger, repo)

	err := service.RenameCategory(ctx, telegramChatID, categoryName, categoryNewName)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestRenameCategoryCategoryDoesNotExists(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().
			RenameCategory(ctx, telegramChatID, categoryName, categoryNewName).
			Times(1).
			Return(core.ErrNotFound),
		logger.EXPECT().Error(core.ErrNotFound).Times(0),
	)
	service := NewCategoryService(logger, repo)

	err := service.RenameCategory(ctx, telegramChatID, categoryName, categoryNewName)
	require.ErrorIs(t, err, core.ErrNotFound)
}

func TestListCategories(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	expected := []core.Category{
		{
			Name: "someCat",
		},
	}
	gomock.InOrder(
		repo.EXPECT().ListCategories(ctx, telegramChatID).Times(1).Return(expected, nil),
		logger.EXPECT().Error(core.ErrNotFound).Times(0),
	)
	service := NewCategoryService(logger, repo)

	result, err := service.ListCategories(ctx, telegramChatID)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestListCategoriesCausedServerError(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	expected := []core.Category{}
	gomock.InOrder(
		repo.EXPECT().ListCategories(ctx, telegramChatID).Times(1).Return(expected, core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewCategoryService(logger, repo)

	result, err := service.ListCategories(ctx, telegramChatID)
	require.ErrorIs(t, err, core.ErrServerError)
	require.Equal(t, expected, result)
}

func TestListCategoriesUserNotFound(t *testing.T) {
	ctx, repo, logger := getCategoryRepositoryAndLogger(t)
	expected := []core.Category{}
	gomock.InOrder(
		repo.EXPECT().ListCategories(ctx, telegramChatID).Times(1).Return(expected, core.ErrNotFound),
		logger.EXPECT().Error(core.ErrNotFound).Times(0),
	)
	service := NewCategoryService(logger, repo)

	result, err := service.ListCategories(ctx, telegramChatID)
	require.ErrorIs(t, err, core.ErrNotFound)
	require.Equal(t, expected, result)
}
