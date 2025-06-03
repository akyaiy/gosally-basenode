package sessions

import "errors"

var (
	ErrSessionIDRequired = errors.New("session ID is required")
	ErrInvalidTTL        = errors.New("invalid TTL value, must be greater than 0")
)
