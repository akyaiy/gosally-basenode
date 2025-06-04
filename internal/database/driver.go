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

	d.driver._internalConnect()
	// switch d.driver.(type) {
	// case *_SQLiteDriver:
	// 	d.Log.Debug("Using SQLite driver for connection")
	// 	if err := d.sqliteConnect(); err != nil {
	// 		d.Log.Error("Failed to connect using SQLite driver", "error", err)
	// 		return err
	// 	}
	// default:
	// 	d.Log.Error("Unsupported driver type for connection", "type", d.driver[DriverTypeSQLite])
	// 	return ErrUnsupportedDriverType
	// }


	// if _, ok := d.driver[DriverTypeSQLite].(*_SQLiteDriver); ok {
	// 	d.Log.Debug("SQLite driver detected")
	// 	d.Log.Debug("Using SQLite driver for connection")
	// 	if err := d.sqliteConnect(); err != nil {
	// 		d.Log.Error("Failed to connect using SQLite driver", "error", err)
	// 	}
	// }

	return nil
}

func (d *Driver) Close() error {
	d.Log.Debug("Closing database connection")

	if _, ok := d.driver[DriverTypeSQLite].(*_SQLiteDriver); ok {
		d.Log.Debug("Closing SQLite driver connection")
		if err := d.sqliteClose(); err != nil {
			d.Log.Error("Failed to close SQLite driver connection", "error", err)
			return err
		}
	}

	return nil
}
