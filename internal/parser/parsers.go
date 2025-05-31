package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func containsKey(slice []string, item string) error {
	for _, s := range slice {
		if s == item {
			return nil
		}
	}
	return errors.New("0x000003e8")
}

func (s *Parser) parseShortParameters(args string, _src config.Source) error {
	logger.CheckAndSetVoid(&s.Log)
	if strings.Contains(args, "=") {
		parts := strings.SplitN(args, "=", 2)
		keys := parts[0]
		valsStr := parts[1]
		vals := strings.Split(valsStr, ",")
		if len(vals) != len(keys) {
			s.Log.Error(fmt.Sprintf("number of keys (%d) does not match number of values (%d)", len(keys), len(vals)))
			return errors.New("0x000003e9")
		}
		for i, k := range keys {
			key := string(k)
			if fullkey, exists := config.Defines.Abbreviations[key]; exists {
				if err := containsKey(config.Defines.Keywords, fullkey); err != nil {
					s.Log.Error("Undefined parameter", "key", key, "err", err)
					continue
				}
				if err := s.setValue(fullkey, vals[i], _src); err != nil {
					s.Log.Error("Failed to set value", "key", key, "err", err)
					continue
				}
			} else {
				s.Log.Error("Undefined parameter", "key", key, "err", "0x000003ea")
				return errors.New("0x000003f2")
			}
		}
		return nil
	}
	er := false
	for _, k := range args {
		key := string(k)
		if fullkey, exists := config.Defines.Abbreviations[key]; exists {
			if err := containsKey(config.Defines.Keywords, fullkey); err != nil {
				s.Log.Error("Undefined parameter", key, "err", err)
				er = true
				continue
			}
			if err := s.setValue(fullkey, config.True, _src); err != nil {
				s.Log.Error("Failed to set value", "key", key, "err", err)
				continue
			}
		} else {
			s.Log.Error("Undefined parameter", "key", key, "err", "0x000003ea")
			er = true
		}
	}
	if !er {
		return nil
	}
	return errors.New("0x000003f2")
}

func (s *Parser) parseLongParameter(arg string, _src config.Source) error {
	if strings.Contains(arg, "=") {
		parts := strings.SplitN(arg, "=", 2)
		key := parts[0]
		val := parts[1]

		if err := containsKey(config.Defines.Keywords, key); err != nil {
			s.Log.Error("Undefined parameter", "key", key, "err", err)
			return errors.New("0x000003f2")
		}
		if err := s.setValue(key, val, _src); err != nil {
			s.Log.Error("Failed to set value", "key", key, "err", err)
			return errors.New("0x000003f2")
		}
		return nil
	}
	if err := containsKey(config.Defines.Keywords, arg); err != nil {
		s.Log.Error("Undefined parameter", "key", arg, "err", err)
		return errors.New("0x000003f2")
	}
	if err := s.setValue(arg, config.True, _src); err != nil {
		s.Log.Error("Failed to set value", "key", arg, "err", err)
		return errors.New("0x000003f2")
	}
	return nil
}
