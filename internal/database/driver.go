package database

// For now, this is a useless driver init implementation.
func (d *Driver) Init() error {
	d.Log.Debug("Initializing database driver")

	return nil
}

func checkConnectionOpts(o *DatabaseConnectionOpt) error {
	if o.ConnectionString == "" {
		return ErrConnectionStringRequired
	}
	if o.Timeout <= 0 {
		return ErrInvalidTimeout
	}
	if o.ConnectionID == "" {
		return ErrConnectionIDRequired
	}
	return nil
}

func (d *Driver) Connect(o *DatabaseConnectionOpt) error {
	if err := checkConnectionOpts(o); err != nil {
		return err
	}
	d.Log.Debug("Connecting to database with connection ID", "id", o.ConnectionID)
	return d.driver._internalConnect(o)
}

func (d *Driver) Close(o *DatabaseConnectionOpt) error {
	if err := checkConnectionOpts(o); err != nil {
		return err
	}
	d.Log.Debug("Closing database connection with connection ID", "id", o.ConnectionID)
	return d.driver._internalClose(o)
}
