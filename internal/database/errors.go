package database

import "errors"

var (
	ErrConnectionIDRequired     = errors.New("connection ID is required")
	ErrConnectionStringRequired = errors.New("connection string is required")
	ErrInvalidTimeout           = errors.New("invalid timeout value, must be greater than 0")
)
