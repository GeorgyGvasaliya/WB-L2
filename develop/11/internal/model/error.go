package model

import "errors"

var (
	ErrNotFound        = errors.New("event not found")
	ErrInvalidInterval = errors.New("interval invalid")
)
