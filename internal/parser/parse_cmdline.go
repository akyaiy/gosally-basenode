package parser

import (
	"errors"
	"os"
	"strings"

	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func (s *Driver) ParseArgs(args []string) error {
	if s.Config.Vars == nil {
		s.Config.Vars = make(config.TypeConfig)
	}
	logger.CheckAndSetVoid(&s.Log)
	var err error
	if err := s.setValue("exec-name", args[0], config.SrcCMD); err != nil {
		s.Log.Error("Failed to set exec-name", "err", err)
		return errors.New("failed to set exec-name")
	}

	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--") {
			err = s.parseLongParameter(arg[2:], config.SrcCMD)
			if err != nil {
				s.Log.Error("Failed to parse cmdline", "err", err)
			}
		} else if strings.HasPrefix(arg, "-") {
			err = s.parseShortParameters(arg[1:], config.SrcCMD)
			if err != nil {
				s.Log.Error("Failed to parse cmdline", "err", err)
			}
		} else {
			err = errors.New("0x000003e8")
			s.Log.Error("Undefined parameter", "key", arg, "err", "0x000003e8")
		}
	}
	if err != nil {
		return errors.New("failed to parse command line arguments")
	}
	return nil
}

func (s *Driver) ParseCMDlineArgs() error {
	return s.ParseArgs(os.Args)
}
