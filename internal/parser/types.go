package parser

import (
	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

type DriverContract interface {
	ParseArgs(args []string) error
	ParseCMDlineArgs() error

	ParseConfigFile(file string) error

	ParseEnvVars() error

	SetConfigValue(_key string, _val string) error
	setValue(key, value string, source config.Source) error

	parseShortParameters(args string, source config.Source) error
	parseLongParameter(arg string, source config.Source) error
}

type DriverBuilderContract interface {
	WithLogger(logger logger.Log) *DriverBuilder
	// Possible to specify configuration, not recommended
	WithConfig(cfg config.Config) *DriverBuilder
	EndBuild() *Driver
}

type DriverBuilder struct {
	driver Driver
}

type Driver struct {
	Log logger.Log
	config.Config
}
