package database

import (
	"database/sql"

	"github.com/akyaiy/gosally-basenode/internal/logger"
	"github.com/akyaiy/gosally-basenode/internal/sessions"
)

const /* Drivers */ (
	DriverTypeSQLite = iota
)

// DriverContract defines the interface for database drivers.
type DriverContract interface {
	// Init initializes the driver and returns a Driver instance.
	Init() error

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

type DriverBuilderContract interface {
	NewDriver() *DriverBuilder
	WithLogger(logger logger.Log) *DriverBuilder
	WithDriverType(driverType int) *DriverBuilder
	EndSafeBuild() (*Driver, error)
	EndBuild() *Driver
}

type DriverBuilder struct {
	driver Driver
}

type Driver struct {
	Log    logger.Log
	driver DriversType
}

// DriversType is a map that holds available database drivers.
type DriversType map[int]any

// _driversDefinitions holds the definitions of available drivers.
var _driversDefinitions DriversType = DriversType{
	DriverTypeSQLite: &_SQLiteDriver{
		SQLite: nil,
	},
	/*
		The following drivers are not implemented yet, but can be added later:
		DriverTypePostgres: &PostgresDriver{
			Postgres: nil,
		},
		etc.
	*/
}

type _SQLiteDriver struct {
	SQLite *sql.DB
}

type ConnectionSessionsDB interface {
	InitSession(o *sessions.Session) error
	QuerySession(o *sessions.Session) (*sessions.Session, error)
	UpdateSession(o *sessions.Session) error
	CloseSession(o *sessions.Session) error
}
