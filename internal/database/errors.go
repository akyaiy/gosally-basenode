package database

import "errors"

var (
	ErrConnectionIDRequired     = errors.New("connection ID is required")
	ErrConnectionStringRequired = errors.New("connection string is required")
	ErrInvalidTimeout           = errors.New("invalid timeout value, must be greater than 0")

	ErrDriverNotFound        = errors.New("driver not found in the available drivers list")
	ErrLoggerNotSet          = errors.New("logger is not set, please set a logger before using the driver")
	ErrUnsupportedDriverType = errors.New("unsupported driver type, please use a supported driver type")
)
