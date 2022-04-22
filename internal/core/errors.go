package core

import "errors"

var (
	ErrNotFound    = errors.New("404 Not Found")
	ErrServerError = errors.New("500 Server Error")

	ErrCategoryExists = errors.New("category already exists")

	ErrWrongCallbackQueryData = errors.New("wrong callback query data")
)
