package lh

import "errors"

var (
	ErrEntryNotFound = errors.New("entry not found")
	ErrUnhashable = errors.New("key not hashable")
)
