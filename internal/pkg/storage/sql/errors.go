package sql

import "errors"

var (
	ErrEntityExists   = errors.New("entity already exists")
	ErrEntityNotFound = errors.New("entity not found")
	ErrForeignKey     = errors.New("foreign key error")
)
