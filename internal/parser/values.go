/*
 * This file contains functions for assigning values ​​to the application config.
 */

package parser

import (
	"errors"

	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func (s *Parser) setValue(_key string, _val string, _src config.Source) error {
	logger.CheckAndSetVoid(&s.Log)
	var (
		tmpSrc  = s.Config.Vars[_key].Source
		rewrite = _src == config.SrcCMD || _src == config.SrcENV
		write   = func() {
			s.Config.Vars[_key] = config.ConfValue{
				Value:  _val,
				Source: _src,
			}
		}
	)
	if tmpSrc == config.SrcNONE {
		write()
		return nil
	} else if tmpSrc == config.SrcCST {
		s.Log.Error("redefining the parameter value is not possible", _key, _val)
		return errors.New("0x0000041a")
	} else if tmpSrc == _src {
		s.Log.Warn("overriding parameter value", _key, _val)
		write()
		return nil
	} else if rewrite {
		s.Log.Warn("overriding parameter value", _key, _val)
		write()
		return nil
	}
	s.Log.Error("redefining the parameter value is not possible", _key, _val)
	return errors.New("0x0000041a")
}
