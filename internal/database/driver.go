package database

// For now, this is a useless driver init implementation.
func (d *Driver) Init() error {
	d.Log.Debug("Initializing database driver")

	return nil
}

func (d *Driver) Connect(o *DatabaseConnectionOpt) error {
	if o.ConnectionString == "" {
		return ErrConnectionStringRequired
	}
	if o.Timeout <= 0 {
		return ErrInvalidTimeout
	}
	if o.ConnectionID == "" {
		return ErrConnectionIDRequired
	}

	d.Log.Debug("Connecting to database with connection ID", "id", o.ConnectionID)

	if _, ok := d.driver[DriverTypeSQLite].(*_SQLiteDriver); ok {
		d.Log.Debug("SQLite driver detected")
		d.Log.Debug("Using SQLite driver for connection")
		if err := d.sqliteConnect(); err != nil {
			d.Log.Error("Failed to connect using SQLite driver", "error", err)
		}
	}

	return nil
}
