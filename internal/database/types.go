package database

import (
	"database/sql"

	"github.com/akyaiy/gosally-basenode/internal/logger"
	"github.com/akyaiy/gosally-basenode/internal/sessions"
)

// DriverContract defines the interface for database drivers.
type DriverContract interface {
	// Init initializes the driver and returns a Driver instance.
	Init() (*Driver, error)

	// Connect connecting to the database using the provided connection string.
	Connect(o *DatabaseConnectionOpt) error

	// Close closes the database connection.
	Close() error
}

type DatabaseConnectionOpt struct {
	ConnectionString string
	Timeout          int64
	ConnectionID     string
}

type ConnectionBuilderContract interface {
	WithConnectionString(connectionString string) *ConnectionBuilder
	WithTimeout(timeout int64) *ConnectionBuilder
	WithConnectionID(connectionID string) *ConnectionBuilder
	EndSafeBuild() (*DatabaseConnectionOpt, error)
	EndBuild() *DatabaseConnectionOpt
}

type ConnectionBuilder struct {
	opts DatabaseConnectionOpt
}

type Driver struct {
	Log    logger.Log
	driver *DriversType
}

// DriversType is a map that holds available database drivers.
type DriversType map[string]any

var Drivers DriversType = DriversType{
	"sqlite": &SQLiteDriver{
		SQLite: nil,
	},
}

type SQLiteDriver struct {
	SQLite *sql.DB
}

type ConnectionSessionsDB interface {
	InitSession(o *sessions.Session) error
	QuerySession(o *sessions.Session) (*sessions.Session, error)
	UpdateSession(o *sessions.Session) error
	CloseSession(o *sessions.Session) error
}
