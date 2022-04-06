package core

import "errors"

var (
	ErrNothingWasDeleted = errors.New("nothing was deleted")
	ErrNothingWasUpdated = errors.New("nothing was updated")
)
