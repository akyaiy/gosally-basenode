package database

import "github.com/akyaiy/gosally-basenode/internal/logger"

func NewConnection() *ConnectionBuilder {
	return &ConnectionBuilder{
		opts: DatabaseConnectionOpt{},
	}
}

func (b *ConnectionBuilder) WithConnectionString(connectionString string) *ConnectionBuilder {
	b.opts.ConnectionString = connectionString
	return b
}

func (b *ConnectionBuilder) WithTimeout(timeout int64) *ConnectionBuilder {
	b.opts.Timeout = timeout
	return b
}

func (b *ConnectionBuilder) WithConnectionID(connectionID string) *ConnectionBuilder {
	b.opts.ConnectionID = connectionID
	return b
}

func (b *ConnectionBuilder) EndSafeBuild() (*DatabaseConnectionOpt, error) {
	if b.opts.ConnectionString == "" {
		return nil, ErrConnectionStringRequired
	}
	if b.opts.Timeout <= 0 {
		return nil, ErrInvalidTimeout
	}
	if b.opts.ConnectionID == "" {
		return nil, ErrConnectionIDRequired
	}
	return &b.opts, nil
}

func (b *ConnectionBuilder) EndBuild() *DatabaseConnectionOpt {
	if b.opts.ConnectionString == "" {
		panic(ErrConnectionStringRequired)
	}
	if b.opts.Timeout <= 0 {
		panic(ErrInvalidTimeout)
	}
	if b.opts.ConnectionID == "" {
		panic(ErrConnectionIDRequired)
	}
	return &b.opts
}

func NewDriver() *DriverBuilder {
	log := logger.NewMockLogger()
	return &DriverBuilder{
		driver: Driver{
			Log: &log, // Logger should be set later
		},
	}
}

func (b *DriverBuilder) WithLogger(log logger.Log) *DriverBuilder {
	b.driver.Log = &log
	return b
}

func (b *DriverBuilder) WithDriverType(driverType int) *DriverBuilder {
	if driver, exists := _driversDefinitions[driverType]; exists {
		b.driver.driver = driver.(DriversType)
		if err := b.driver.driver.SetLogger(b.driver.Log); err != nil {
			panic(err) // Ensure the driver has a logger set
		} // Ensure the driver has a logger set
		// Set the logger for the driver
	} else {
		panic(ErrDriverNotFound)
	}
	return b
}

func (b *DriverBuilder) EndSafeBuild() (*Driver, error) {
	if b.driver.Log.Logger == nil {
		return nil, ErrLoggerNotSet
	}

	if b.driver.driver == nil {
		return nil, ErrDriverNotFound
	}
	return &b.driver, nil
}

func (b *DriverBuilder) EndBuild() *Driver {
	if b.driver.Log.Logger == nil {
		panic(ErrLoggerNotSet)
	}

	if b.driver.driver == nil {
		panic(ErrDriverNotFound)
	}
	return &b.driver
}
