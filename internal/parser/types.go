package parser

import (
	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

type ConfigParse interface {
	ParseArgs(args []string) error
	ParseCMDlineArgs() error

	ParseConfigFile(file string) error

	ParseEnvVars() error

	SetConfigValue(_key string, _val string) error
	setValue(key, value string, source config.Source) error

	parseShortParameters(args string, source config.Source) error
	parseLongParameter(arg string, source config.Source) error
}

type Parser struct {
	Log logger.Log
	config.Config
}
