package errors

import "errors"

var (
	ErrAlreadyExists = errors.New("resource alteady exists")
	ErrInvalidInput  = errors.New("invalid input provided")
	ErrNotFound      = errors.New("resource not found")
	ErrUnauthorized  = errors.New("unauthorized access")
)