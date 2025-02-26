package analyze

import (
	"errors"
)

var (
	ErrAnalyzeExists   = errors.New("analyze account already exists")
	ErrAnalyzeNotFound = errors.New("analyze account not found")
	ErrForeignKeyError = errors.New("foreign key error")
	ErrInvalidField    = errors.New("error invalid field")
)
