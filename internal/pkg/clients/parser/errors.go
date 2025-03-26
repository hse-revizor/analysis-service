package parser

import "errors"

var (
	ErrRequest                 = errors.New("Request error")
	ErrBadURL                  = errors.New("Bad URL")
	ErrUnauthorized            = errors.New("Unauthorized")
	ErrNotEnoughRights         = errors.New("Not enough rights")
	ErrBadRequest              = errors.New("Bad request")
	ErrInternal                = errors.New("Internal error")
	ErrReachedMaxRequestNumber = errors.New("Maximum request count")
	ErrInvalidEntity           = errors.New("Invalid entity")
	ErrUnknown                 = errors.New("Unknown API Error")
	ErrNotFound                = errors.New("Entity not found")
)
