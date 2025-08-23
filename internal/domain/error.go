package domain

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrEmptyField         = errors.New("field is empty")
	ErrMessageSentRefused = errors.New("message sent refused")
)
