package store

import (
	"errors"
)

var (
	ErrDuplicate = errors.New("Duplicate value")
	ErrNotFound  = errors.New("Value not found")
	ErrInternal  = errors.New("Internal error occured")
)
