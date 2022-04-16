package core

import "errors"

var (
	ErrNotFound          = errors.New("404 not found")
	ErrNothingWasDeleted = errors.New("nothing was deleted")
	ErrNothingWasUpdated = errors.New("nothing was updated")

	ErrCategoryExists = errors.New("category already exists")

	ErrWrongCallbackQueryData = errors.New("wrong callback query data")

	// telegram message errors
	ErrNotAuthorized = errors.New("You are not authorized! Type /start command to authorize") // todo config?
)
