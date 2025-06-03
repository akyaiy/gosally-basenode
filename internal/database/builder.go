package database

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
