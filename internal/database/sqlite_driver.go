package database

import (
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

// "database/sql"
// "errors"
// "github.com/akyaiy/GoSally-node/internal/logger"
// _ "modernc.org/sqlite"
// "os"
// "path/filepath"
// "strings"

func (d *_SQLiteDriver) SetLogger(log *logger.Log) error {
	if log == nil {
		return ErrLoggerNotSet
	}
	d.Log = log
	return nil
}

func (d *_SQLiteDriver) GetLogger() (*logger.Log, error) {
	if d.Log == nil {
		return nil, ErrLoggerNotSet
	}
	return d.Log, nil
}

func (d *_SQLiteDriver) _internalConnect(o *DatabaseConnectionOpt) error {
	d.Log.Debug("Connecting to SQLite database", "connectionID", o.ConnectionID, "timeout", o.Timeout)
	// just let me cook
	return nil
}

func (d *_SQLiteDriver) _internalClose(o *DatabaseConnectionOpt) error {
	// just let me cook
	return nil
}
