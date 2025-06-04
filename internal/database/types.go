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

type DriversType _internalDriverContract

// _driversDefinitions holds the definitions of available drivers.
var _driversDefinitions map[int]any = map[int]any{
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

type _internalDriverContract interface {
	_internalConnect() error
	_internalClose() error
}

type _SQLiteDriver struct {
	SQLite *sql.DB
}

type SessionStorageContract interface {
	SaveSession(o *sessions.Session) error
	QuerySession(o *sessions.Session) (*sessions.Session, error)
	UpdateSession(o *sessions.Session) error
	DeleteSession(o *sessions.Session) error
}

type SessionStorage struct {
	Log logger.Log
	//driver 
}