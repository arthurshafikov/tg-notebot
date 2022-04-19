package core

import "errors"

var (
	ErrNotFound          = errors.New("404 not found")
	ErrServerError       = errors.New("500 server error")
	ErrNothingWasDeleted = errors.New("nothing was deleted")
	ErrNothingWasUpdated = errors.New("nothing was updated")

	ErrCategoryExists = errors.New("category already exists")

	ErrWrongCallbackQueryData = errors.New("wrong callback query data")
)
