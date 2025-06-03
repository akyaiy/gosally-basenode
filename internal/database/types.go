package database

import (
	"database/sql"

	"github.com/akyaiy/gosally-basenode/internal/logger"
	"github.com/akyaiy/gosally-basenode/internal/sessions"
)

type DBDriver interface {
	Init() (*Driver, error)

	Connect(o string) error
	Close() error
}

type Driver struct {
	Log    logger.Log
	driver *DriversAvailable
}

type DriversAvailable struct {
	SQLite *sql.DB
}

type ConnectionSessionsDB interface {
	InitSession(o *sessions.Session) error
	QuerySession(o *sessions.Session) (*sessions.Session, error)
	UpdateSession(o *sessions.Session) error
	CloseSession(o *sessions.Session) error
}
