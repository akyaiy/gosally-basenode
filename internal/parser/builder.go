package parser

import (
	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func NewDriver() *DriverBuilder {
	driver := &DriverBuilder{
		driver: Driver{
			Config: config.Config{},
			Log:    *logger.NewMockLogger(),
		},
	}

	driver.driver.Config.Vars = make(config.TypeConfig)
	return driver
}

func (b *DriverBuilder) WithLogger(log logger.Log) *DriverBuilder {
	b.driver.Log = log
	return b
}

func (b *DriverBuilder) WithConfig(cfg config.Config) *DriverBuilder {
	b.driver.Config = cfg
	return b
}

func (b *DriverBuilder) EndBuild() *Driver {
	return &b.driver
}
