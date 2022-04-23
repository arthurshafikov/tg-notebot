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

var noteContent = "someNoteContent"

func getNoteRepositoryAndLogger(
	t *testing.T,
) (context.Context, *mock_repository.MockNotes, *mock_services.MockLogger) {
	t.Helper()
	ctrl := gomock.NewController(t)

	return context.Background(), mock_repository.NewMockNotes(ctrl), mock_services.NewMockLogger(ctrl)
}

func TestAddNote(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().AddNote(ctx, telegramChatID, categoryName, noteContent).Times(1).Return(nil),
	)
	service := NewNoteService(logger, repo)

	err := service.AddNote(ctx, telegramChatID, categoryName, noteContent)
	require.NoError(t, err)
}

func TestAddNoteReturnError(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().AddNote(ctx, telegramChatID, categoryName, noteContent).Times(1).Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewNoteService(logger, repo)

	err := service.AddNote(ctx, telegramChatID, categoryName, noteContent)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestRemoveNote(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().
			RemoveNote(ctx, telegramChatID, categoryName, noteContent).
			Times(1).
			Return(nil),
		logger.EXPECT().Error(core.ErrServerError).Times(0),
	)
	service := NewNoteService(logger, repo)

	err := service.RemoveNote(ctx, telegramChatID, categoryName, noteContent)
	require.NoError(t, err)
}

func TestRemoveNoteReturnServerError(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().
			RemoveNote(ctx, telegramChatID, categoryName, noteContent).
			Times(1).
			Return(core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewNoteService(logger, repo)

	err := service.RemoveNote(ctx, telegramChatID, categoryName, noteContent)
	require.ErrorIs(t, err, core.ErrServerError)
}

func TestRemoveNoteReturnNotFound(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	gomock.InOrder(
		repo.EXPECT().
			RemoveNote(ctx, telegramChatID, categoryName, noteContent).
			Times(1).
			Return(core.ErrNotFound),
		logger.EXPECT().Error(core.ErrNotFound).Times(0),
	)
	service := NewNoteService(logger, repo)

	err := service.RemoveNote(ctx, telegramChatID, categoryName, noteContent)
	require.ErrorIs(t, err, core.ErrNotFound)
}

func TestListNotesFromCategory(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	expected := []core.Note{
		{
			Content: "someNote",
		},
	}
	gomock.InOrder(
		repo.EXPECT().ListNotesFromCategory(ctx, telegramChatID, categoryName).Times(1).Return(expected, nil),
	)
	service := NewNoteService(logger, repo)

	notes, err := service.ListNotesFromCategory(ctx, telegramChatID, categoryName)
	require.NoError(t, err)
	require.Equal(t, expected, notes)
}

func TestListNotesFromCategoryReturnServerError(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	expected := []core.Note{}
	gomock.InOrder(
		repo.EXPECT().
			ListNotesFromCategory(ctx, telegramChatID, categoryName).
			Times(1).
			Return(expected, core.ErrServerError),
		logger.EXPECT().Error(core.ErrServerError).Times(1),
	)
	service := NewNoteService(logger, repo)

	notes, err := service.ListNotesFromCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrServerError)
	require.Equal(t, expected, notes)
}

func TestListNotesFromCategoryUserOrCategoryNotFound(t *testing.T) {
	ctx, repo, logger := getNoteRepositoryAndLogger(t)
	expected := []core.Note{}
	gomock.InOrder(
		repo.EXPECT().
			ListNotesFromCategory(ctx, telegramChatID, categoryName).
			Times(1).
			Return(expected, core.ErrNotFound),
	)
	service := NewNoteService(logger, repo)

	notes, err := service.ListNotesFromCategory(ctx, telegramChatID, categoryName)
	require.ErrorIs(t, err, core.ErrNotFound)
	require.Equal(t, expected, notes)
}
